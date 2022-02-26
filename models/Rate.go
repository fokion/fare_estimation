package models

type Rate struct {
	Price       float64
	IsIdle      bool
	StartHour   int8
	StartMinute int8
	EndHour     int8
	EndMinute   int8
}
