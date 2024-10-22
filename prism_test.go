package gogeometry_test

import (
	"testing"

	gogeometry "github.com/hamonangann/go-geometry"

	"github.com/stretchr/testify/assert"
)

func TestPrism(t *testing.T) {
	t.Run("should failed constructing Prism if height is negative", func(t *testing.T) {
		base, _ := gogeometry.NewSquare(1.0)
		_, err := gogeometry.NewPrism(base, -5.0)
		assert.ErrorContains(t, err, "cannot be negative")
	})

	t.Run("should construct Prism if height is non-negative", func(t *testing.T) {
		base, _ := gogeometry.NewSquare(1.0)
		p, _ := gogeometry.NewPrism(base, 5.0)
		assert.IsType(t, &gogeometry.Prism{}, p)
	})

	t.Run("should failed construct cylinder as Prism if radius is negative", func(t *testing.T) {
		_, err := gogeometry.NewCylinder(-1.0, 5.0)
		assert.ErrorContains(t, err, "cannot be negative")
	})

	t.Run("should failed construct cylinder as Prism if height is negative", func(t *testing.T) {
		_, err := gogeometry.NewCylinder(1.0, -5.0)
		assert.ErrorContains(t, err, "cannot be negative")
	})

	t.Run("should construct cylinder as Prism if both radius and height is non-negative", func(t *testing.T) {
		p, _ := gogeometry.NewCylinder(1.0, 5.0)
		assert.IsType(t, &gogeometry.Prism{}, p)
	})

	t.Run("should return base correctly", func(t *testing.T) {
		base, _ := gogeometry.NewSquare(1.0)
		p, _ := gogeometry.NewPrism(base, 2.0)
		assert.Equal(t, base, p.Base())
	})

	t.Run("should return height correctly", func(t *testing.T) {
		base, _ := gogeometry.NewSquare(1.0)
		p, _ := gogeometry.NewPrism(base, 2.0)
		assert.Equal(t, 2.0, p.Height())
	})

	t.Run("should return surface area correctly", func(t *testing.T) {
		base, _ := gogeometry.NewSquare(1.0)
		p, _ := gogeometry.NewPrism(base, 5.0)
		assert.Equal(t, 22.0, p.SurfaceArea())
	})

	t.Run("should return volume correctly", func(t *testing.T) {
		base, _ := gogeometry.NewSquare(1.0)
		p, _ := gogeometry.NewPrism(base, 5.0)
		assert.Equal(t, 5.0, p.Volume())
	})
}
