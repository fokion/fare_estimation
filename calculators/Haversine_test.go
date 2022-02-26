package calculators

import (
	"fare_estimation/errors"
	"fare_estimation/models"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestCalculate_with_no_points_so_it_throws_error(t *testing.T) {
	haversine := NewHarversineCalculatorInKM()
	_, err := haversine.GetDistance(nil, nil)
	if assert.Error(t, err) {
		assert.Equal(t, errors.CalculationError{Message: "Missing both points so I cannot calculate any distance"}, err)
	} else {
		t.Fail()
	}

}
func TestCalculate_with_no_from_point_so_it_throws_error(t *testing.T) {
	haversine := NewHarversineCalculatorInKM()
	_, err := haversine.GetDistance(nil, &models.Point{Latitude: 55.93985, Longitude: -3.22046})
	if assert.Error(t, err) {
		assert.Equal(t, errors.CalculationError{Message: "Missing From point so I cannot calculate any distance"}, err)
	} else {
		t.Fail()
	}

}
func TestCalculate_with_no_to_points_so_it_throws_error(t *testing.T) {
	haversine := NewHarversineCalculatorInKM()
	_, err := haversine.GetDistance(&models.Point{Latitude: 55.94429, Longitude: -3.20623}, nil)
	if assert.Error(t, err) {
		assert.Equal(t, errors.CalculationError{Message: "Missing To point so I cannot calculate any distance"}, err)
	} else {
		t.Fail()
	}

}
func TestCalculate_in_a_straight_line(t *testing.T) {
	//given two points 55.94429 -3.20623 and 55.93985  -3.22046 it is a ~1.0137 km line
	haversine := NewHarversineCalculatorInKM()
	distance, err := haversine.GetDistance(&models.Point{Latitude: 55.94429, Longitude: -3.20623},
		&models.Point{Latitude: 55.93985, Longitude: -3.22046})

	if err == nil {
		assert.Equal(t, new(big.Float).SetPrec(4).SetFloat64(1.0137), new(big.Float).SetPrec(4).SetFloat64(distance))
	} else {
		t.Fail()
	}

}
