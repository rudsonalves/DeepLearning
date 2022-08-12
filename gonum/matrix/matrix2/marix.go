package main

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

func main() {
	a := mat.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})
	b := mat.NewDense(3, 3, []float64{8, 9, 10, 1, 4, 2, 9, 0, 2})

	c := mat.NewDense(3, 2, []float64{3, 2, 1, 4, 0, 8})

	// Add a and b
	d := mat.NewDense(3, 3, nil)
	d.Add(a, b)
	fd := mat.Formatted(d)
	fmt.Printf("d = a + b = \n%0.4v\n\n", fd)

	// Multply a and c
	f := mat.NewDense(3, 2, nil)
	f.Mul(a, c)
	ff := mat.Formatted(f)
	fmt.Println("f = a * c:\n", ff)

	// Raising a matrix to a power
	g := mat.NewDense(3, 3, nil)
	g.Pow(a, 5)
	fg := mat.Formatted(g)
	fmt.Println("g = a^5:\n", fg)

	// Apply a function to each of the elements of a
	h := mat.NewDense(3, 3, nil)
	sqrt := func(_, _ int, v float64) float64 { return math.Sqrt(v) }
	h.Apply(sqrt, a)
	fh := mat.Formatted(h)
	fmt.Println("h = sqrt(a):\n", fh)
}
