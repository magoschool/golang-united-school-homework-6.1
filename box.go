package golang_united_school_homework

import (
	"errors"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

var (
	ErrMaxCapacity    = errors.New("max capacity reached")
	ErrInvalidIndex   = errors.New("index out of bounds")
	ErrCircleNotFound = errors.New("circle not found")
)

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) >= b.shapesCapacity {
		return ErrMaxCapacity
	}

	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i < 0 || i >= len(b.shapes) {
		return nil, ErrInvalidIndex
	}

	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i < 0 || i >= len(b.shapes) {
		return nil, ErrInvalidIndex
	}

	lShape := b.shapes[i]
	lNewShapes := make([]Shape, 0, len(b.shapes)-1)
	for j := 0; j < i; j++ {
		lNewShapes = append(lNewShapes, b.shapes[j])
	}
	for j := i + 1; j < len(b.shapes); j++ {
		lNewShapes = append(lNewShapes, b.shapes[j])
	}

	b.shapes = lNewShapes
	return lShape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i < 0 || i >= len(b.shapes) {
		return nil, ErrInvalidIndex
	}

	lShape := b.shapes[i]
	b.shapes[i] = shape

	return lShape, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var lSum float64 = 0

	for _, lShape := range b.shapes {
		lSum = lSum + lShape.CalcPerimeter()
	}

	return lSum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var lSum float64 = 0

	for _, lShape := range b.shapes {
		lSum = lSum + lShape.CalcArea()
	}

	return lSum
}

func isCircle(aShape Shape) bool {
	_, lOk := aShape.(*Circle)
	return lOk
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var lCount = 0
	for _, lShape := range b.shapes {
		if isCircle(lShape) {
			lCount++
		}
	}

	if lCount == 0 {
		return ErrCircleNotFound
	}

	lNewShapes := make([]Shape, 0, len(b.shapes)-lCount)
	for i := 0; i < len(b.shapes); i++ {
		lShape := b.shapes[i]
		if !isCircle(lShape) {
			lNewShapes = append(lNewShapes, lShape)
		}
	}

	b.shapes = lNewShapes
	return nil
}
