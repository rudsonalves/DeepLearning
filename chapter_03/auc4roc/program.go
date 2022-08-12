package main

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/integrate"
	"gonum.org/v1/gonum/stat"
)

func main() {

	// Define our scores and classes.
	scores := []float64{8, 7.5, 6, 5, 3, 0}
	classes := []bool{true, true, true, false, true, true}
	weights := []float64{2, 2, 3, 6, 1, 4}
	n := 9

	// Define our scores and classes.
	// scores := []float64{0.1, 0.35, 0.4, 0.8}
	// classes := []bool{true, false, true, false}
	// weights := []float64{1, 1, 1, 1}
	// n := 5

	// Calculate the true positive rates (recalls) and
	// false positive rates.
	stat.SortWeightedLabeled(scores, classes, weights)
	cutoffs := make([]float64, n)
	floats.Span(cutoffs, math.Nextafter(scores[0], scores[0]-1), scores[len(scores)-1])

	tpr, fpr, _ := stat.ROC(cutoffs, scores, classes, weights)

	// Compute the Area Under Curve.
	auc := integrate.Trapezoidal(fpr, tpr)

	// Output the results to standard out.
	fmt.Printf("true  positive rate: %.3f\n", tpr)
	fmt.Printf("false positive rate: %.3f\n", fpr)
	fmt.Printf("auc: %.3f\n", auc)
}
