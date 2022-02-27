package models

type Journey struct {
	ID     string
	sum    float64
	Points []*Point
}

type Trip interface {
	GetID() string
	GetPoints() []*Point
	GetTotalFare() float64
	SetTotalFare(fare float64)
	Add(fare float64)
}

func (t *Journey) Add(fare float64) {
	t.sum += fare
}
func (t *Journey) GetPoints() []*Point {
	return t.Points
}

func (t *Journey) GetTotalFare() float64 {
	return t.sum
}

func (t *Journey) SetTotalFare(fare float64) {
	t.sum = fare
}
