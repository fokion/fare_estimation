package calculators

import (
	"fare_estimation/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculate_in_a_straight_line(t *testing.T) {
	//given two points 55.94429 -3.20623 and 55.93985  -3.22046 it is a ~0.9656 km line
	haversine := Haversine{
		From: &models.Point{Latitude: 55.94429, Longitude: -3.20623},
		To:   &models.Point{Latitude: 55.93985, Longitude: -3.22046},
	}
	distance, err := haversine.GetDistance()
	if err != nil {
		assert.Equal(t, 0.9656, distance)
	}
	t.Fail()

}
