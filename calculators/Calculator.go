package calculators

import "math"

type Calculator interface {
	GetDistance() float64
}

func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

const (
	EARTH_RADIUS_IN_KM    = 6367
	EARTH_RADIUS_IN_MILES = 3958
)
