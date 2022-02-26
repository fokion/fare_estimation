package models

type Rate struct {
	Price       float64
	IsIdle      bool
	StartHour   int8
	StartMinute int8
	EndHour     int8
	EndMinute   int8
}

const (
	FLAG_RATE    = 1.3
	MINIMUM_RATE = 3.47
)
