# Go-Geometry

Simple Go package to perform operations on planar and spatial objects.

### Installation

```shell
go get github.com/hamonangann/go-geometry
```

### Example

```go
package main

import (
	"fmt"

	gogeometry "github.com/hamonangann/go-geometry"
)

func main() {
	var sq gogeometry.Plane
	sq, _ = gogeometry.NewSquare(5)
	fmt.Printf("Perimeter of the square is: %.2f\n", sq.Perimeter())
	fmt.Printf("Area of the square is: %.2f\n", sq.Area())
}

```