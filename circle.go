package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (c Circle) CalcPerimeter() float64 {
	// CalcPerimeter returns calculation result of perimeter
	return 2 * math.Pi * c.Radius
}

func (c Circle) CalcArea() float64 {
	// CalcArea returns calculation result of area
	return math.Pi * c.Radius * c.Radius
}
