package gogeometry

import (
	"fmt"
	"math"
)

type Sphere struct {
	radius float64
}

// SurfaceArea implements Space.
func (s Sphere) SurfaceArea() float64 {
	return 4 * math.Pi * s.radius * s.radius
}

// Volume implements Space.
func (s Sphere) Volume() float64 {
	return (4 * math.Pi * s.radius * s.radius * s.radius) / 3
}

func (s Sphere) Radius() float64 {
	return s.radius
}

func NewSphere(radius float64) (*Sphere, error) {
	if radius < 0 {
		return nil, fmt.Errorf("sphere radius cannot be negative")
	}

	return &Sphere{radius}, nil
}
