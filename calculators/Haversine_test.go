package calculators

import "testing"

func TestCalculate_in_a_straight_line(t *testing.T) {
	//given two points 55.94429 -3.20623 and 55.93985  -3.22046 it is a ~0.9656 km line
	assert.Equal(t, 0.9656)
}
