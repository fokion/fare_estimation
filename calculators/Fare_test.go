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
	farecalc := NewFareCalculator(NewHarversineCalculatorInKM(), NewSpeedCalculatorInKM(), GetDefaultRates(), 10, models.FLAG_RATE, models.MINIMUM_RATE)

	date := time.Date(2022, 02, 25, 10, 0, 0, 0, time.UTC)
	fare, err := farecalc.CalculateFare(&models.Point{Latitude: 55.94429, Longitude: -3.20623, Timestamp: date.Unix()},
		&models.Point{Latitude: 55.93985, Longitude: -3.22046, Timestamp: date.Add(time.Minute * 5).Unix()})
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, new(big.Float).SetPrec(4).SetFloat64(0.7501), new(big.Float).SetPrec(4).SetFloat64(fare))

}

func TestCalculateFare_with_idle_speed(t *testing.T) {
	farecalc := NewFareCalculator(NewHarversineCalculatorInKM(), NewSpeedCalculatorInKM(), GetDefaultRates(), 10, models.FLAG_RATE, models.MINIMUM_RATE)

	date := time.Date(2022, 02, 25, 10, 0, 0, 0, time.UTC)
	fare, err := farecalc.CalculateFare(&models.Point{Latitude: 55.94429, Longitude: -3.20623, Timestamp: date.Unix()},
		&models.Point{Latitude: 55.93985, Longitude: -3.22046, Timestamp: date.Add(time.Minute * 12).Unix()})
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, new(big.Float).SetPrec(4).SetFloat64(IDLE_RATE*12/60), new(big.Float).SetPrec(4).SetFloat64(fare))

}

func TestCalculateFare_with_more_than_max_speed(t *testing.T) {
	farecalc := NewFareCalculator(NewHarversineCalculatorInKM(), NewSpeedCalculatorInKM(), GetDefaultRates(), 10, models.FLAG_RATE, models.MINIMUM_RATE)

	date := time.Date(2022, 02, 25, 10, 0, 0, 0, time.UTC)
	_, err := farecalc.CalculateFare(&models.Point{Latitude: 55.94429, Longitude: -3.20623, Timestamp: date.Unix()},
		&models.Point{Latitude: 55.93985, Longitude: -3.22046, Timestamp: date.Add(time.Second * 1).Unix()})
	assert.Error(t, err)
}

func TestCleanup(t *testing.T) {
	farecalc := NewFareCalculator(NewHarversineCalculatorInKM(), NewSpeedCalculatorInKM(), GetDefaultRates(), 10, models.FLAG_RATE, models.MINIMUM_RATE)

	date := time.Date(2022, 02, 25, 10, 0, 0, 0, time.UTC)

	points := []*models.Point{}
	points = append(points, &models.Point{Latitude: 55.94429, Longitude: -3.20623, Timestamp: date.Unix()})
	points = append(points, &models.Point{Latitude: 55.93985, Longitude: -3.22046, Timestamp: date.Add(time.Second * 1).Unix()})

	points = append(points, &models.Point{Latitude: 55.93985, Longitude: -3.22046, Timestamp: date.Add(time.Minute * 10).Unix()})
	cleanPoints := farecalc.CleanUpPoints(points)

	assert.Equal(t, 2, len(cleanPoints))
	assert.Equal(t, date.Add(time.Minute*10).Unix(), cleanPoints[len(cleanPoints)-1].Timestamp)
}

func TestCalculateFare_when_we_are_changing_rates(t *testing.T) {

	distanceCalculator := NewHarversineCalculatorInKM()
	speedCalculator := NewSpeedCalculatorInKM()
	farecalc := NewFareCalculator(distanceCalculator, speedCalculator, GetDefaultRates(), 10, models.FLAG_RATE, models.MINIMUM_RATE)
	date := time.Date(2022, 02, 25, 23, 54, 0, 0, time.UTC)

	pointA := &models.Point{Latitude: 55.94429, Longitude: -3.20623, Timestamp: date.Unix()}
	pointB := &models.Point{Latitude: 55.93985, Longitude: -3.22046, Timestamp: date.Add(time.Minute * 5).Unix()}
	pointC := &models.Point{Latitude: 55.94429, Longitude: -3.20623, Timestamp: date.Add(time.Minute * 10).Unix()}

	_, err := farecalc.CalculateFare(pointA, pointB)
	if err != nil {
		t.Fail()
	}
	//this is after the switch so there will be a portion with the old fare and some with the new one
	fareReturn, err := farecalc.CalculateFare(pointB, pointC)

	//speed in point B to C
	distanceCovered, _ := distanceCalculator.GetDistance(pointB, pointC)
	speed, _ := speedCalculator.GetSpeed(distanceCovered, pointB.Timestamp, pointC.Timestamp)

	assert.True(t, speed < speedCalculator.GetMaxSpeed())

	//the difference between time in point B and going to the next rate is 1 minute

	distanceWithinDayRate := float64(speed) / 60.0
	expected := distanceWithinDayRate*DAY_RATE + (distanceCovered-distanceWithinDayRate)*NIGHT_RATE

	assert.Equal(t, new(big.Float).SetPrec(4).SetFloat64(expected), new(big.Float).SetPrec(4).SetFloat64(fareReturn))

}
