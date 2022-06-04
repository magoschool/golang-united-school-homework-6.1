package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (t Triangle) CalcPerimeter() float64 {
	// CalcPerimeter returns calculation result of perimeter
	return 3 * t.Side
}

func (t Triangle) CalcArea() float64 {
	// CalcArea returns calculation result of area
	return t.Side * t.Side * math.Sqrt(3) / 4
}
