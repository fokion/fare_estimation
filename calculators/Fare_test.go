package calculators

import (
	"github.com/stretchr/testify/assert"
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
