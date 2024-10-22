package gogeometry

import (
	"fmt"
	"math"
)

type Triangle struct {
	sideA float64
	sideB float64
	sideC float64
}

// Area implements Plane.
func (t Triangle) Area() float64 {
	s := 0.5 * t.Perimeter()
	return math.Sqrt(s * (s - t.sideA) * (s - t.sideB) * (s - t.sideC))
}

// Perimeter implements Plane.
func (t Triangle) Perimeter() float64 {
	return t.sideA + t.sideB + t.sideC
}

func (t Triangle) Sides() (float64, float64, float64) {
	return t.sideA, t.sideB, t.sideC
}

func NewTriangle(sideA float64, sideB float64, sideC float64) (*Triangle, error) {
	if sideA < 0 || sideB < 0 || sideC < 0 {
		return nil, fmt.Errorf("triangle sides cannot be negative")
	}

	return &Triangle{sideA, sideB, sideC}, nil
}

func NewEquilateralTriangle(side float64) (*Triangle, error) {
	return NewTriangle(side, side, side)
}
