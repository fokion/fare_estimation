package calculators

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIsPartOfTimeRange(t *testing.T) {

	//time 17:05 is between 17:00 and 17:15
	currentTime := time.Date(2022, 02, 25, 17, 5, 0, 0, time.UTC)
	//time 17:05 is between 17:00 and 17:15
	assert.True(t, IsPartOfTimeRange(currentTime, 17, 0, 17, 15), "time 17:05 is between 17:00 and 17:15")

	//time 17:05 is between 17:00 and 00:01
	assert.True(t, IsPartOfTimeRange(currentTime, 17, 0, 0, 1), "time 17:05 is between 17:00 and 00:01")

	//time 17:05 is between 17:00 and 17:06
	assert.True(t, IsPartOfTimeRange(currentTime, 17, 0, 17, 6), "time 17:05 is between 17:00 and 17:06")
	//time 17:05 is NOT between 17:00 and 17:04
	assert.False(t, IsPartOfTimeRange(currentTime, 17, 0, 17, 4), "time 17:05 is NOT between 17:00 and 17:04")

	currentTime = time.Date(2022, 02, 25, 16, 59, 0, 0, time.UTC)
	assert.False(t, IsPartOfTimeRange(currentTime, 17, 0, 0, 1), "time 16:59 is NOT between 17:00 and 0:01")
}
