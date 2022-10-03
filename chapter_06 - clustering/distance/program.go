package main

import (
	"fmt"

	"gonum.org/v1/gonum/floats"
)

func main() {
	distance := floats.Distance([]float64{0, 0}, []float64{3, 4}, 3)

	fmt.Printf("Distance: %.2f\n", distance)
}
