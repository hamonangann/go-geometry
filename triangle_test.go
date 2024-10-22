package gogeometry_test

import (
	"testing"

	gogeometry "github.com/hamonangann/go-geometry"

	"github.com/stretchr/testify/assert"
)

func TestTriangle(t *testing.T) {
	t.Run("should failed constructing Triangle if one of the sides is negative", func(t *testing.T) {
		_, err := gogeometry.NewTriangle(2.0, 4.0, -1.0)
		assert.ErrorContains(t, err, "cannot be negative")
	})

	t.Run("should construct Triangle if all of the sides is non-negative", func(t *testing.T) {
		tr, _ := gogeometry.NewTriangle(2.0, 4.0, 1.0)
		assert.IsType(t, &gogeometry.Triangle{}, tr)
	})

	t.Run("should construct equilateral Triangle if the side is non-negative", func(t *testing.T) {
		tr, _ := gogeometry.NewEquilateralTriangle(1.0)
		assert.IsType(t, &gogeometry.Triangle{}, tr)
	})

	t.Run("should return sides correctly", func(t *testing.T) {
		tr, _ := gogeometry.NewTriangle(2.0, 4.0, 1.0)
		sideA, sideB, sideC := tr.Sides()

		assert.Equal(t, 2.0, sideA)
		assert.Equal(t, 4.0, sideB)
		assert.Equal(t, 1.0, sideC)
	})

	t.Run("should return perimeter correctly", func(t *testing.T) {
		tr, _ := gogeometry.NewTriangle(2.0, 4.0, 1.0)
		assert.Equal(t, 7.0, tr.Perimeter())
	})

	t.Run("should return area correctly", func(t *testing.T) {
		tr, _ := gogeometry.NewTriangle(6.0, 5.0, 5.0)
		assert.Equal(t, 12.0, tr.Area())
	})
}
