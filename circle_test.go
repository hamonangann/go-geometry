package gogeometry_test

import (
	"testing"

	gogeometry "github.com/hamonangann/go-geometry"

	"github.com/stretchr/testify/assert"
)

func TestCircle(t *testing.T) {
	t.Run("should failed constructing Circle if radius is negative", func(t *testing.T) {
		_, err := gogeometry.NewCircle(-1.0)
		assert.ErrorContains(t, err, "cannot be negative")
	})

	t.Run("should construct Circle if radius is non-negative", func(t *testing.T) {
		c, _ := gogeometry.NewCircle(1.0)
		assert.IsType(t, &gogeometry.Circle{}, c)
	})

	t.Run("should return radius correctly", func(t *testing.T) {
		c, _ := gogeometry.NewCircle(1.0)
		assert.Equal(t, 1.0, c.Radius())
	})

	t.Run("should return perimeter correctly", func(t *testing.T) {
		c, _ := gogeometry.NewCircle(1.0)
		assert.InDelta(t, 6.28, c.Perimeter(), 0.1)
	})

	t.Run("should return area correctly", func(t *testing.T) {
		c, _ := gogeometry.NewCircle(1.0)
		assert.InDelta(t, 3.14, c.Area(), 0.1)
	})
}
