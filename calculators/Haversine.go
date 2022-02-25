package calculators

import (
	"fare_estimation/errors"
	"fare_estimation/models"
)

type Haversine struct {
	From *models.Point
	To   *models.Point
}

func (calculator *Haversine) GetDistance() (float64, error) {
	if calculator.From != nil && calculator.To != nil {

	} else if calculator.To != nil {
		return 0.0, errors.CalculationError{Message: "Missing From point so I cannot calculate any distance"}
	} else if calculator.From != nil {
		return 0.0, errors.CalculationError{Message: "Missing To point so I cannot calculate any distance"}
	}
	return 0.0, errors.CalculationError{Message: "Missing both points so I cannot calculate any distance"}
}
