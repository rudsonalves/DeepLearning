package main

import (
	"fmt"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

func main() {
	v1 := []float64{1.0, 1.0, 1.0}
	v2 := []float64{5.0, 3.8, 0.0}

	fmt.Println(floats.Dot(v1, v2))
	fmt.Println(floats.Norm(v1, 2))
	floats.Add(v1, v2)
	fmt.Println(v1)

	va := mat.NewVecDense(3, []float64{1.0, 1.0, 1.0})
	vb := mat.NewVecDense(3, []float64{5.0, 3.8, 0.0})

	fmt.Println(mat.Dot(va, vb))
	fmt.Println(mat.Norm(va, 2))
	//fmt.Println(va + vb)
	fmt.Println(va)

}
