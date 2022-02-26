package calculators

import (
	"fare_estimation/models"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
	"time"
)

func TestGetMovingRateForTimestamp(t *testing.T) {
	rates := GetDefaultRates()

	//10:00 in the morning
	timestamp := time.Date(2022, 02, 25, 10, 0, 0, 0, time.UTC)

	assert.Equal(t, DAY_RATE, GetMovingRateForTimestamp(rates, timestamp.Unix()).Price, "10:00 in the morning must have 0.74 per km")

	timestamp = time.Date(2022, 02, 25, 05, 0, 0, 0, time.UTC)

	assert.Equal(t, NIGHT_RATE, GetMovingRateForTimestamp(rates, timestamp.Unix()).Price, "5:00 has night rate")

	timestamp = time.Date(2022, 02, 25, 04, 58, 0, 0, time.UTC)

	assert.Equal(t, NIGHT_RATE, GetMovingRateForTimestamp(rates, timestamp.Unix()).Price, "4:58 has night rate")

	timestamp = time.Date(2022, 02, 25, 0, 1, 0, 0, time.UTC)

	assert.Equal(t, NIGHT_RATE, GetMovingRateForTimestamp(rates, timestamp.Unix()).Price, "00:01 has the night rate")

	timestamp = time.Date(2022, 02, 25, 23, 59, 0, 0, time.UTC)

	assert.Equal(t, DAY_RATE, GetMovingRateForTimestamp(rates, timestamp.Unix()).Price, "23:59 has the day rate")

	timestamp = time.Date(2022, 02, 25, 0, 0, 0, 0, time.UTC)

	assert.Equal(t, DAY_RATE, GetMovingRateForTimestamp(rates, timestamp.Unix()).Price, "0:0 has the day rate")

}

func TestCalculateFare_with_moving_speed(t *testing.T) {
	farecalc := NewFareCalculator(NewHarversineCalculatorInKM(), NewSpeedCalculatorInKM(), GetDefaultRates(), 10)

	date := time.Date(2022, 02, 25, 10, 0, 0, 0, time.UTC)
	fare, err := farecalc.CalculateFare(&models.Point{Latitude: 55.94429, Longitude: -3.20623, Timestamp: date.Unix()},
		&models.Point{Latitude: 55.93985, Longitude: -3.22046, Timestamp: date.Add(time.Minute * 5).Unix()})
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, new(big.Float).SetPrec(4).SetFloat64(0.7501), new(big.Float).SetPrec(4).SetFloat64(fare))

}

func TestCalculateFare_with_idle_speed(t *testing.T) {
	farecalc := NewFareCalculator(NewHarversineCalculatorInKM(), NewSpeedCalculatorInKM(), GetDefaultRates(), 10)

	date := time.Date(2022, 02, 25, 10, 0, 0, 0, time.UTC)
	fare, err := farecalc.CalculateFare(&models.Point{Latitude: 55.94429, Longitude: -3.20623, Timestamp: date.Unix()},
		&models.Point{Latitude: 55.93985, Longitude: -3.22046, Timestamp: date.Add(time.Minute * 12).Unix()})
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, new(big.Float).SetPrec(4).SetFloat64(IDLE_RATE*12/60), new(big.Float).SetPrec(4).SetFloat64(fare))

}

func TestCalculateFare_with_more_than_max_speed(t *testing.T) {
	farecalc := NewFareCalculator(NewHarversineCalculatorInKM(), NewSpeedCalculatorInKM(), GetDefaultRates(), 10)

	date := time.Date(2022, 02, 25, 10, 0, 0, 0, time.UTC)
	_, err := farecalc.CalculateFare(&models.Point{Latitude: 55.94429, Longitude: -3.20623, Timestamp: date.Unix()},
		&models.Point{Latitude: 55.93985, Longitude: -3.22046, Timestamp: date.Add(time.Second * 1).Unix()})
	assert.Error(t, err)
}
