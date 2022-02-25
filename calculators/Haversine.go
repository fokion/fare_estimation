package calculators

import (
	"fare_estimation/errors"
	"fare_estimation/models"
	"math"
)

type Haversine struct {
	From *models.Point
	To   *models.Point
}

// GetDistance calculates the distance in kilometers or returns an error with what is missing
func (calculator *Haversine) GetDistance() (float64, error) {
	if calculator.From != nil && calculator.To != nil {
		fromLatitudeInRadians := degreesToRadians(calculator.From.Latitude)
		fromLongitudeInRadians := degreesToRadians(calculator.From.Longitude)
		toLatitudeInRadians := degreesToRadians(calculator.To.Latitude)
		toLongitudeInRadians := degreesToRadians(calculator.To.Longitude)
		deltaLat := toLatitudeInRadians - fromLatitudeInRadians
		deltaLon := toLongitudeInRadians - fromLongitudeInRadians

		a := math.Pow(math.Sin(deltaLat/2), 2) + math.Cos(fromLatitudeInRadians)*math.Cos(toLatitudeInRadians)*
			math.Pow(math.Sin(deltaLon/2), 2)

		circleDistanceInRadians := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
		return EARTH_RADIUS_IN_KM * circleDistanceInRadians, nil

	} else if calculator.To != nil {
		return 0.0, errors.CalculationError{Message: "Missing From point so I cannot calculate any distance"}
	} else if calculator.From != nil {
		return 0.0, errors.CalculationError{Message: "Missing To point so I cannot calculate any distance"}
	}
	return 0.0, errors.CalculationError{Message: "Missing both points so I cannot calculate any distance"}
}
