package gogeometry_test

import (
	"testing"

	gogeometry "github.com/hamonangann/go-geometry"

	"github.com/stretchr/testify/assert"
)

func TestRectangle(t *testing.T) {
	t.Run("should failed constructing Rectangle if widthis negative", func(t *testing.T) {
		_, err := gogeometry.NewRectangle(-1.0, 2.0)
		assert.ErrorContains(t, err, "cannot be negative")
	})

	t.Run("should failed constructing Rectangle if length is negative", func(t *testing.T) {
		_, err := gogeometry.NewRectangle(1.0, -2.0)
		assert.ErrorContains(t, err, "cannot be negative")
	})

	t.Run("should construct Rectangle if both width and length is non-negative", func(t *testing.T) {
		r, _ := gogeometry.NewRectangle(1.0, 2.0)
		assert.IsType(t, &gogeometry.Rectangle{}, r)
	})

	t.Run("should construct square as Rectangle if the side is non-negative", func(t *testing.T) {
		r, _ := gogeometry.NewSquare(1.0)
		assert.IsType(t, &gogeometry.Rectangle{}, r)
	})

	t.Run("should return width correctly", func(t *testing.T) {
		r, _ := gogeometry.NewRectangle(1.0, 2.0)
		assert.Equal(t, 1.0, r.Width())
	})

	t.Run("should return length correctly", func(t *testing.T) {
		r, _ := gogeometry.NewRectangle(1.0, 2.0)
		assert.Equal(t, 2.0, r.Length())
	})

	t.Run("should return perimeter correctly", func(t *testing.T) {
		r, _ := gogeometry.NewRectangle(1.0, 2.0)
		assert.Equal(t, 6.0, r.Perimeter())
	})

	t.Run("should return area correctly", func(t *testing.T) {
		r, _ := gogeometry.NewRectangle(1.0, 2.0)
		assert.Equal(t, 2.0, r.Area())
	})
}
