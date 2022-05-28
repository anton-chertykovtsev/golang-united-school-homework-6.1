package golang_united_school_homework

import (
	"errors"
	"fmt"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

var (
	ErrOutOfRange = errors.New("out of range")
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
	if len(b.shapes)+1 > b.shapesCapacity {
		return fmt.Errorf("AddShape %w", ErrOutOfRange)
	}

	b.shapes = append(b.shapes, shape)

	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i > len(b.shapes)-1 {
		return nil, fmt.Errorf("GetByIndex %w", ErrOutOfRange)
	}

	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i > len(b.shapes)-1 {
		return nil, fmt.Errorf("ExtractByIndex %w", ErrOutOfRange)
	}

	s := b.shapes[i]

	empty := NewBox(len(b.shapes) - 1)
	for k, v := range b.shapes {
		if k != i {
			empty.AddShape(v)
		}
	}
	b.shapes = empty.shapes

	return s, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i > len(b.shapes)-1 {
		return nil, fmt.Errorf("ExtractByIndex %w", ErrOutOfRange)
	}

	s := b.shapes[i]

	empty := NewBox(len(b.shapes))
	for k, v := range b.shapes {
		if k != i {
			empty.AddShape(v)
		} else {
			empty.AddShape(shape)
		}
	}
	b.shapes = empty.shapes

	return s, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64

	for _, v := range b.shapes {
		sum = sum + v.CalcPerimeter()
	}

	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64

	for _, v := range b.shapes {
		sum = sum + v.CalcArea()
	}

	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {

	empty := NewBox(len(b.shapes))
	for _, v := range b.shapes {
		switch v.(type) {
		case *Circle:
		default:
			_ = empty.AddShape(v)
		}
	}
	if len(b.shapes) == len(empty.shapes) {
		return errors.New("circles not founded")
	} else {
		b.shapes = empty.shapes
	}

	return nil

}
