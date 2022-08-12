package main

import (
	"fmt"
	"log"

	"gonum.org/v1/gonum/mat"
)

func main() {
	a := mat.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})

	// Compute and output the trsnapose of the matrix
	at := a.T()
	ft := mat.Formatted(at, mat.Prefix(" "))
	fmt.Printf("a^T = %v\n\n", ft)

	// Compute and output the determinant of a
	detA := mat.Det(a)
	fmt.Printf("det(a) = %.2f\n\n", detA)

	// Compute and output the inverse of a
	aInverse := mat.NewDense(3, 3, nil)
	if err := aInverse.Inverse(a); err != nil {
		log.Fatal(err)
	}
	fi := mat.Formatted(aInverse, mat.Prefix(" "))
	fmt.Println("a^-1:\n", fi)
}
