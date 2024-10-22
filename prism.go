package gogeometry

import "fmt"

type Prism struct {
	base   Plane
	height float64
}

// SurfaceArea implements Space.
func (p Prism) SurfaceArea() float64 {
	return (2 * p.base.Area()) + (p.base.Perimeter() * p.height)

}

// Volume implements Space.
func (p Prism) Volume() float64 {
	return p.base.Area() * p.height
}

func (p Prism) Base() Plane {
	return p.base
}

func (p Prism) Height() float64 {
	return p.height
}

func NewPrism(base Plane, height float64) (*Prism, error) {
	if height < 0 {
		return nil, fmt.Errorf("prism height cannot be negative")
	}

	return &Prism{base, height}, nil
}

func NewCylinder(radius float64, height float64) (*Prism, error) {
	base, err := NewCircle(radius)
	if err != nil {
		return nil, err
	}

	return NewPrism(base, height)
}
