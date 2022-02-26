package main

import (
	"bufio"
	"fare_estimation/calculators"
	"fare_estimation/converters"
	"fare_estimation/models"
	"flag"
	"fmt"
	"github.com/pkg/profile"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	fileName          string
	outputFilePath    string
	defaultOutputName string
)

func main() {

	os.Exit(realMain())
}

func realMain() int {
	//format the date as YY-MM-DD_hh_mm
	defaultOutputName = fmt.Sprintf("output_%s.txt", time.Now().Format("06-01-02_15_04"))
	flag.StringVar(&fileName, "f", "", "Specify a file to parse")
	flag.StringVar(&outputFilePath, "o", defaultOutputName, fmt.Sprintf("Specify an output file that will be relative to the execution path . Default is %s", defaultOutputName))

	flag.Parse()
	defer profile.Start(profile.CPUProfile).Stop()

	// after declaring flags we need to call it

	if strings.Trim(fileName, " ") == "" {
		fmt.Println("You have not specified a file to parse you need to use the flag -f with the path of the file")
		return 1
	}

	//make it relative to execution
	outputFilePath = fmt.Sprintf("./%s", outputFilePath)

	waitgroup := sync.WaitGroup{}

	lineChannel := make(chan string, 2)

	journeyChannel := make(chan *models.Journey, 5)

	fareCalculator := calculators.NewFareCalculator(
		calculators.NewHarversineCalculatorInKM(),
		calculators.NewSpeedCalculatorInKM(),
		calculators.GetDefaultRates(),
		10.0,
		models.FLAG_RATE,
		models.MINIMUM_RATE,
	)

	waitgroup.Add(3)
	//file reader routine
	go func(ch chan string, wg *sync.WaitGroup, filePath string) {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println(err)
			fmt.Println(fmt.Sprintf("could not open file %s", filePath))
			panic("could not open the file")
		}
		//close file
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				panic("could not close the file")
			}
		}(file)

		//close channel in the end
		defer close(ch)

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ch <- scanner.Text()
		}

		wg.Done()
	}(lineChannel, &waitgroup, fileName)
	//receive the feed from lines
	go func(lineCh chan string, journeyChannel chan *models.Journey, wg *sync.WaitGroup) {

		var cache = map[string]*models.Journey{}
		prevKey := ""
		//receive values until the channel is closed and the buffer queue is
		//empty
		for line := range lineCh {
			id, point, err := converters.ConvertLineToPoint(line)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if prevKey != "" && id != prevKey {
				completedJourney, _ := cache[prevKey]
				journeyChannel <- completedJourney
				delete(cache, prevKey)
			}
			journey, containsKey := cache[id]
			if !containsKey {
				prevKey = id
				cache[id] = &models.Journey{ID: id, Points: []*models.Point{}}
				journey, _ = cache[id]
			}
			list := journey.Points
			journey.Points = append(list, point)
			cache[id] = journey
		}
		for id := range cache {
			completedJourney, _ := cache[id]
			journeyChannel <- completedJourney
			delete(cache, prevKey)
		}
		close(journeyChannel)
		wg.Done()
	}(lineChannel, journeyChannel, &waitgroup)
	//fare calculation for a completed journey
	go func(journeyChannel chan *models.Journey, calc calculators.FareCalculator, wg *sync.WaitGroup, outputFile string) {
		for journey := range journeyChannel {
			allPoints := calc.CleanUpPoints(journey.GetPoints())
			//initialise the fare with the flag rate as someone entered the taxi
			journey.SetTotalFare(calc.GetFlagRate())
			startingPoint := allPoints[0]
			for i := 1; i < len(allPoints); i++ {
				point := allPoints[i]
				fare, err := calc.CalculateFare(startingPoint, point)
				if err != nil {
					fmt.Println(fmt.Sprintf("ride='%s' %s", journey.ID, err))
				} else {

					journey.Add(fare)
					startingPoint = point
				}
			}
			if journey.GetTotalFare() < calc.GetMinimumRate() {
				journey.SetTotalFare(calc.GetMinimumRate())
			}

			f, err := os.OpenFile(outputFile,
				os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}

			if _, err := f.WriteString(fmt.Sprintf("%s,%f\n", journey.ID, journey.GetTotalFare())); err != nil {
				log.Println(err)
			}
			f.Close()
			fmt.Println(fmt.Sprintf("ride='%s' fare=%f", journey.ID, journey.GetTotalFare()))

		}
		wg.Done()
	}(journeyChannel, fareCalculator, &waitgroup, outputFilePath)

	waitgroup.Wait()
	return 0
}
