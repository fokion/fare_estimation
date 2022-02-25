package calculators

import "fare_estimation/errors"

//GetSpeedInKm calculate the speed based on the distance covered.
func GetSpeedInKm(distanceInKm float64, startTimestamp int64, endTimestamp int64) (float64, error) {
	deltaTime := (endTimestamp - startTimestamp) / 3600
	if deltaTime > 0 {
		return distanceInKm / float64(deltaTime), nil
	}
	return 0, errors.CalculationError{
		Message: "The time is not greater than 0",
	}
}
