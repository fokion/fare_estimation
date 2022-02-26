package calculators

import (
	"fare_estimation/errors"
	"fare_estimation/models"
	"fmt"
	"math"
)

const (
	DAY_RATE   = 0.74
	NIGHT_RATE = 1.3
	IDLE_RATE  = 11.9
)

type FareHandler struct {
	IdleSpeed          float64
	Rates              []*models.Rate
	SpeedCalculator    SpeedCalculator
	DistanceCalculator DistanceCalculator
	minimumRate        float64
	flagRate           float64
}

func (f *FareHandler) CalculateFare(from *models.Point, to *models.Point) (float64, error) {
	distance, err := f.DistanceCalculator.GetDistance(from, to)
	if err == nil {
		speed, speedErr := f.SpeedCalculator.GetSpeed(distance, from.Timestamp, to.Timestamp)
		if speedErr != nil {
			return 0, speedErr
		}
		idleRate := GetIdleRate(f.GetRates())
		if speed > f.SpeedCalculator.GetMaxSpeed() {
			return 0, errors.CalculationError{
				Message: fmt.Sprintf("More than %s so skipping", f.SpeedCalculator.PrintMaxSpeed()),
			}
		} else if speed <= math.Abs(f.IdleSpeed) {
			return float64(to.Timestamp-from.Timestamp) * idleRate.Price / 3600, nil
		}
		//the assumption is that the timestamps are really close to each other, so we just need to
		//check if the rate in the timestamp at the start is the same with the rate in the end
		//if not then calculate the distance based on the speed for the duration and then the rest is calculated
		//with the other rate
		startingRate := GetMovingRateForTimestamp(f.Rates, from.Timestamp)
		endRate := GetMovingRateForTimestamp(f.Rates, to.Timestamp)
		if startingRate.Price == endRate.Price {
			return distance * startingRate.Price, nil
		}
		//we need to calculate based on the distance
		duration := GetTimeDifference(GetTimeFromTimestamp(from.Timestamp), int(startingRate.EndHour), int(startingRate.EndMinute), 0)
		coveredDistance := speed * duration.Hours()
		return coveredDistance*startingRate.Price + (distance-coveredDistance)*endRate.Price, nil

	}
	return 0, err
}

func (f *FareHandler) GetRates() []*models.Rate {
	return f.Rates
}

func (f *FareHandler) GetMinimumRate() float64 {
	return f.minimumRate
}
func (f *FareHandler) GetFlagRate() float64 {
	return f.flagRate
}

type FareCalculator interface {
	CalculateFare(from *models.Point, to *models.Point) (float64, error)
	GetRates() []*models.Rate
	GetMinimumRate() float64
	GetFlagRate() float64
}

func NewFareCalculator(distanceCalculator DistanceCalculator, speedCalculator SpeedCalculator, rates []*models.Rate, idleSpeed float64, flagRate float64, minimumRate float64) FareCalculator {

	if rates == nil {
		rates = GetDefaultRates()
	}
	return &FareHandler{
		DistanceCalculator: distanceCalculator,
		SpeedCalculator:    speedCalculator,
		Rates:              rates,
		IdleSpeed:          idleSpeed,
		flagRate:           flagRate,
		minimumRate:        minimumRate,
	}
}

func GetDefaultRates() []*models.Rate {
	var rates []*models.Rate

	rates = append(rates, &models.Rate{Price: DAY_RATE, StartHour: 5, StartMinute: 1, IsIdle: false})
	rates = append(rates, &models.Rate{Price: NIGHT_RATE, StartMinute: 1, EndHour: 5, IsIdle: false})
	rates = append(rates, &models.Rate{Price: IDLE_RATE, IsIdle: true})
	return rates
}

func GetIdleRate(rates []*models.Rate) *models.Rate {
	for _, rate := range rates {
		if rate.IsIdle {
			return rate
		}
	}
	return &models.Rate{Price: IDLE_RATE, IsIdle: true}
}

func GetMovingRateForTimestamp(rates []*models.Rate, timestamp int64) *models.Rate {
	for _, rate := range rates {
		if !rate.IsIdle {
			currentTime := GetTimeFromTimestamp(timestamp)
			if IsPartOfTimeRange(currentTime, int(rate.StartHour), int(rate.StartMinute), int(rate.EndHour), int(rate.EndMinute)) {
				return rate
			}
		}
	}
	return nil
}
