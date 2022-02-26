package converters

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertLineToPoint(t *testing.T) {
	id, point, err := ConvertLineToPoint("9,37.965566,23.733313,1405587958")
	if err != nil {
		assert.Fail(t, fmt.Sprintf(" Failed as it threw the following error `%s` ", err))
	}

	assert.Equal(t, "9", id)

	assert.Equal(t, 37.965566, point.Latitude)
	assert.Equal(t, 23.733313, point.Longitude)
	assert.Equal(t, int64(1405587958), point.Timestamp)
}

func TestConvertLineToPoint_trim_on_sides(t *testing.T) {
	id, point, err := ConvertLineToPoint("    9,37.965566,23.733313,1405587958   ")
	if err != nil {
		assert.Fail(t, fmt.Sprintf(" Failed as it threw the following error `%s` ", err))
	}

	assert.Equal(t, "9", id)

	assert.Equal(t, 37.965566, point.Latitude)
	assert.Equal(t, 23.733313, point.Longitude)
	assert.Equal(t, int64(1405587958), point.Timestamp)
}
func TestConvertLineToPoint_trim_on_among_values(t *testing.T) {
	id, point, err := ConvertLineToPoint("9 , 37.965566 , 23.733313       , 1405587958   ")
	if err != nil {
		assert.Fail(t, fmt.Sprintf(" Failed as it threw the following error `%s` ", err))
	}

	assert.Equal(t, "9", id)

	assert.Equal(t, 37.965566, point.Latitude)
	assert.Equal(t, 23.733313, point.Longitude)
	assert.Equal(t, int64(1405587958), point.Timestamp)
}
