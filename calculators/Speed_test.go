package calculators

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetSpeedInKm(t *testing.T) {
	timeNow := time.Now()
	timeReached := timeNow.Add(time.Hour * 1)
	calc := SpeedInKMCalculator{}
	expected, _ := calc.GetSpeed(100, timeNow.Unix(), timeReached.Unix())
	assert.Equal(t, float64(100), expected)
}

func TestGetSpeedInKm_test_20(t *testing.T) {
	timeNow := time.Now()
	timeReached := timeNow.Add(time.Hour * 1)
	calc := SpeedInKMCalculator{}
	expected, _ := calc.GetSpeed(20, timeNow.Unix(), timeReached.Unix())
	assert.Equal(t, float64(20), expected)
}

func TestGetSpeedInKm_test_error(t *testing.T) {
	timeNow := time.Now()
	calc := SpeedInKMCalculator{}
	_, err := calc.GetSpeed(20, timeNow.Unix(), timeNow.Unix())
	assert.Error(t, err)
}

func TestSpeedInSmallerScale(t *testing.T) {
	date := time.Date(2022, 02, 25, 10, 0, 0, 0, time.UTC)
	calc := SpeedInKMCalculator{}
	expected, _ := calc.GetSpeed(1, date.Unix(), date.Add(time.Minute*4).Unix())
	assert.Equal(t, float64(15), expected)
}
