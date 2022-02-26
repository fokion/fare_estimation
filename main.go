package main

import (
	"bufio"
	"fare_estimation/calculators"
	"fare_estimation/converters"
	"fare_estimation/models"
	"fmt"
	"os"
	"sync"
)

func main() {

	waitgroup := sync.WaitGroup{}

	lineChannel := make(chan string, 2)

	journeyChannel := make(chan *models.Journey, 5)

	fileName := "./paths.csv"

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
		defer wg.Done()

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
	}(lineChannel, journeyChannel, &waitgroup)
	//fare calculation for a completed journey
	go func(journeyChannel chan *models.Journey, calc calculators.FareCalculator, wg *sync.WaitGroup) {
		defer wg.Done()
		for journey := range journeyChannel {
			allPoints := journey.GetPoints()
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
			fmt.Println(fmt.Sprintf("ride='%s' fare=%f", journey.ID, journey.GetTotalFare()))
		}
	}(journeyChannel, fareCalculator, &waitgroup)

	waitgroup.Wait()
}
