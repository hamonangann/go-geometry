package gogeometry

import (
	"fmt"
)

type Rectangle struct {
	width  float64
	length float64
}

// Area implements Plane.
func (r Rectangle) Area() float64 {
	return r.width * r.length
}

// Perimeter implements Plane.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.length)
}

func (r Rectangle) Width() float64 {
	return r.width
}

func (r Rectangle) Length() float64 {
	return r.length
}

func NewRectangle(width float64, length float64) (*Rectangle, error) {
	if width < 0 {
		return nil, fmt.Errorf("rectangle width cannot be negative")
	}

	if length < 0 {
		return nil, fmt.Errorf("rectangle length cannot be negative")
	}

	return &Rectangle{width, length}, nil
}

func NewSquare(side float64) (*Rectangle, error) {
	return NewRectangle(side, side)
}
