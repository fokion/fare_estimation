package calculators

import (
	"fare_estimation/models"
	"math"
)

type DistanceCalculator interface {
	GetDistance(from *models.Point, to *models.Point) (float64, error)
}

func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

const (
	EARTH_RADIUS_IN_KM    = 6367
	EARTH_RADIUS_IN_MILES = 3958
)
