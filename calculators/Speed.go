package calculators

import (
	"fare_estimation/errors"
	"fmt"
)

const (
	MAX_SPEED_IN_KM = 100
)

type SpeedCalculator interface {
	GetSpeed(distance float64, startTimestamp int64, endTimestamp int64) (float64, error)
	GetMaxSpeed() float64
	PrintMaxSpeed() string
}

type SpeedInKMCalculator struct {
	Max float64
}

func NewSpeedCalculatorInKM() *SpeedInKMCalculator {
	return &SpeedInKMCalculator{
		MAX_SPEED_IN_KM,
	}
}

func (c *SpeedInKMCalculator) GetMaxSpeed() float64 {
	return c.Max
}

func (c *SpeedInKMCalculator) PrintMaxSpeed() string {
	return fmt.Sprintf("%fkm/h", c.Max)
}

//GetSpeed calculate the speed based on the distance covered.
func (c *SpeedInKMCalculator) GetSpeed(distanceInKm float64, startTimestamp int64, endTimestamp int64) (float64, error) {
	deltaTime := (endTimestamp - startTimestamp) / 3600
	if deltaTime > 0 {
		return distanceInKm / float64(deltaTime), nil
	}
	return 0, errors.CalculationError{
		Message: "The time is not greater than 0",
	}
}
