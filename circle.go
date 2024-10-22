package gogeometry

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

// Area implements Plane.
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Perimeter implements Plane.
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) Radius() float64 {
	return c.radius
}

func NewCircle(radius float64) (*Circle, error) {
	if radius < 0 {
		return nil, fmt.Errorf("circle radius cannot be negative")
	}
	return &Circle{radius}, nil
}
