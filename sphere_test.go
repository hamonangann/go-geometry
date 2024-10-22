package gogeometry_test

import (
	"testing"

	gogeometry "github.com/hamonangann/go-geometry"

	"github.com/stretchr/testify/assert"
)

func TestSphere(t *testing.T) {
	t.Run("should failed constructing Sphere if radius is negative", func(t *testing.T) {
		_, err := gogeometry.NewSphere(-5.0)
		assert.ErrorContains(t, err, "cannot be negative")
	})

	t.Run("should construct Sphere if radius is non-negative", func(t *testing.T) {
		s, _ := gogeometry.NewSphere(5.0)
		assert.IsType(t, &gogeometry.Sphere{}, s)
	})

	t.Run("should return radius correctly", func(t *testing.T) {
		s, _ := gogeometry.NewSphere(5.0)
		assert.Equal(t, 5.0, s.Radius())
	})

	t.Run("should return surface area correctly", func(t *testing.T) {
		s, _ := gogeometry.NewSphere(1.0)
		assert.InDelta(t, 12.56, s.SurfaceArea(), 0.1)
	})

	t.Run("should return volume correctly", func(t *testing.T) {
		s, _ := gogeometry.NewSphere(1.0)
		assert.InDelta(t, 4.18, s.Volume(), 0.1)
	})
}
