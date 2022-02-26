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
