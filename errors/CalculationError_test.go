package errors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculationError_Error_formatting(t *testing.T) {
	err := CalculationError{
		Message: "test",
	}
	assert.Equal(t, "type=\"Calculation\" details=\"test\"", err.Error())
}
