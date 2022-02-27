# Fare Calculator

Started working from the calculators and the models towards the main method that has the go routines.

## Calculators
### Distance

Created an interface called `DistanceCalculator` that is implemented by the `Haversine` where you can define the prefered `radius` 
that you can use in order to do get the calculation in the appropriate metric. 

I had seen a pattern due to the lack of constructors that you can have methods that start with `New` that generate
an instance of a `struct`. As a result I have created the method `NewHarversineCalculatorInKM` that initialises it with the
Earth's radius.

### Speed

The speed calculator has a similar approach with an interface , with distance as an argument and the two timestamps. In that struct
I have added the maximum allowed speed in order to be able to filter.


### Time 

In that file I have some helper methods that are used to get the time difference between a timestamp and a time in the ranges `GetTimeDifference`.
Another helper function is `IsPartOfTimeRange` which checks if a timestamp is within the ranges based on the hour , minute provided .

### Fare

This was the last file that I've worked on in the calculators and holds a `FareHandler` that implements the `FareCalculator` interface
that has the following self explaining methods `CleanUpPoints`, `CalculateFare` , `GetRates` , `GetMinimumRate` and lastly `GetFlagRate` which 
are used in calculation. We have here some helper functions such as the `GetIdleRate` from a list of rates with a default being returned if 
nothing is found and the `GetMovingRateForTimestamp` which gives the rate that fits a time range. The `CalculateFare` does a calculation 
between two provided points and throws an error if something is not right either in the distance or speed calculation.



### Models

## Point

`Point` is a struct which holds the coordinates and a timestamp of it

## Rate

This is a struct where it holds fields related to the starting and ending hour and minute as well as the price and if the rate
is for the vehicle staying idle. In that file there are some constants such as the minimum speed , flag rate and the minimum rate that 
will be used when we do the fare calculation


## Trip

`Trip` is an interface which is implemented by `Journey` which holds the `ID` , array of `Points` and a `sum` .


## Putting all together


We have 2 channels one for the lines of the file and one for Journeys. We initialise the fare calculator, we create a 
waitgroup of 3 as we will have one go routine to parse the file line by line and send that in the channel where another go routine
converts the line into a Point ( see `LineParser` ) and holds them in a map that removes the entry when it can see another journey starting up.
The last go routine is using the fare calculator to add the `flag` fare ,  iterate the points and check if the fare in the end is below the 
minimum amount in order to set it to it. Last but not least opens a file to append that line and writes in the console.


---

- `f` this is a mandatory flag as it defines the location of the file that we have to parse
- `o` this is an optional flag which defaults to the following pattern `output_YY-MM-DD_hh_mm.txt`
---

  Here is an example of how to run the project
 
```shell
go build
./fare_estimation.exe -f .\paths.csv -o test.txt
```

## Comments & Improvements 

This was my first time writing a Go program from scratch after a while , some things are different compared to other programming 
languages that I have used the most notable for me were the lack of constructors and how we declare the interfaces. In the fare calculation i had 
the speed filter embedded as I thought that this will make it easier for not needing the cleanup but after reading the instructions it was 
required to have those journeys cleaned up in advance. One other thing that I have misread in the beginning was how the journeys are 
defined in the file as initially thought that it can be mixed thus the reason why I left a map structure there in the go routine .


