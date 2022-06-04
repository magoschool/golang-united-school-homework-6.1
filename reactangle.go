package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (r Rectangle) CalcPerimeter() float64 {
	// CalcPerimeter returns calculation result of perimeter
	return 2 * (r.Height + r.Weight)
}

func (r Rectangle) CalcArea() float64 {
	// CalcArea returns calculation result of area
	return r.Height * r.Weight
}
