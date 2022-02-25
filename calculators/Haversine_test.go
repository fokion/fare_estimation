package calculators

import (
	"fare_estimation/errors"
	"fare_estimation/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculate_with_no_points_so_it_throws_error(t *testing.T) {
	haversine := Haversine{}
	_, err := haversine.GetDistance()
	if assert.Error(t, err) {
		assert.Equal(t, errors.CalculationError{Message: "Missing both points so I cannot calculate any distance"}, err)
	} else {
		t.Fail()
	}

}
func TestCalculate_with_no_from_point_so_it_throws_error(t *testing.T) {
	haversine := Haversine{To: &models.Point{Latitude: 55.93985, Longitude: -3.22046}}
	_, err := haversine.GetDistance()
	if assert.Error(t, err) {
		assert.Equal(t, errors.CalculationError{Message: "Missing From point so I cannot calculate any distance"}, err)
	} else {
		t.Fail()
	}

}
func TestCalculate_with_no_to_points_so_it_throws_error(t *testing.T) {
	haversine := Haversine{From: &models.Point{Latitude: 55.94429, Longitude: -3.20623}}
	_, err := haversine.GetDistance()
	if assert.Error(t, err) {
		assert.Equal(t, errors.CalculationError{Message: "Missing To point so I cannot calculate any distance"}, err)
	} else {
		t.Fail()
	}

}
func TestCalculate_in_a_straight_line(t *testing.T) {
	//given two points 55.94429 -3.20623 and 55.93985  -3.22046 it is a ~0.9656 km line
	haversine := Haversine{
		From: &models.Point{Latitude: 55.94429, Longitude: -3.20623},
		To:   &models.Point{Latitude: 55.93985, Longitude: -3.22046},
	}
	distance, err := haversine.GetDistance()
	if err != nil {
		assert.Equal(t, 0.9656, distance)
	} else {
		t.Fail()
	}

}
