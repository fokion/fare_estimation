package errors

import "fmt"

type CalculationError struct {
	Message string
}

func (C CalculationError) Error() string {
	return fmt.Sprintf("type=%s details=%s", "Calculation", C.Message)
}
