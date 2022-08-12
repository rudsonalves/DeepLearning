package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
)

func main() {
	f, err := os.Open("iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	irisDF := dataframe.ReadCSV(f)

	sepalLength := irisDF.Col("petal_length").Float()

	minVal := floats.Min(sepalLength)
	maxVal := floats.Max(sepalLength)
	rangeVal := maxVal - minVal
	varianceVal := stat.Variance(sepalLength, nil)
	stdDevVal := stat.StdDev(sepalLength, nil)

	// Sort the values
	inds := make([]int, len(sepalLength))
	floats.Argsort(sepalLength, inds)

	// Get the Quantiles
	quant25 := stat.Quantile(.25, stat.Empirical, sepalLength, nil)
	quant50 := stat.Quantile(.50, stat.Empirical, sepalLength, nil)
	quant75 := stat.Quantile(.75, stat.Empirical, sepalLength, nil)

	// Output the results to standard out.
	fmt.Printf("\nSepal Length Summary Statistics:\n")
	fmt.Printf("Max value: %0.2f\n", maxVal)
	fmt.Printf("Min value: %0.2f\n", minVal)
	fmt.Printf("Range value: %0.2f\n", rangeVal)
	fmt.Printf("Variance value: %0.2f\n", varianceVal)
	fmt.Printf("Std Dev value: %0.2f\n", stdDevVal)
	fmt.Printf("25 Quantile: %0.2f\n", quant25)
	fmt.Printf("50 Quantile: %0.2f\n", quant50)
	fmt.Printf("75 Quantile: %0.2f\n\n", quant75)
	fmt.Println(sepalLength)
}
