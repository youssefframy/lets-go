package circle

import "fmt"

const (
	PI = 3.14
)

type Circle struct {
	radius float64
}

func NewCircle(radius float64) Circle {
	return Circle{
		radius: radius,
	}
}

func (c Circle) CalculateCircumference() {
	circumference := 2 * PI * c.radius
	fmt.Println(circumference)
}