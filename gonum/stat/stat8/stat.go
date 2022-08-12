package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
)

/*

Next, we need to calculate the p-value corresponding to this Chi-square. This requires us to know information about the hi-square distribution, which defines the p-values for certain meseasures of chi-square and certain degrees of freedom.

*/

func main() {
	// Define the observed frequencies.
	observed := []float64{
		260.0, // This number is the number of observed with no regular exercise.
		135.0, // This number is the number of observed with sporatic exercise.
		105.0, // This number is the number of observed with regular exercise.
	}

	// Define the total observed.
	totalObserved := 500.0

	// Calculate the expected frequencies (again assuming the null Hypothesis).
	expected := []float64{
		totalObserved * 0.60,
		totalObserved * 0.25,
		totalObserved * 0.15,
	}

	// Calculate the ChiSquare test statistic.
	chiSquare := stat.ChiSquare(observed, expected)

	// Output the test statistic to standard out.
	fmt.Println("Observed  Expected   Diff")
	for i := 0; i < len(observed); i++ {
		fmt.Printf(" %3.0f       %3.0f      %3.0f\n", observed[i], expected[i], observed[i]-expected[i])
	}
	fmt.Printf("\nChi-square: %0.2f\n", chiSquare)

	// create a Chi-square distribution with K degrees of freedom.
	// In this case we have K = 3-1 = 2, because degrees of freedom
	// for a Chi-square distribution is the number of possible
	// categories minus one.
	chiDist := distuv.ChiSquared{
		K:   2.0,
		Src: nil,
	}

	// Calculate the p-value for our specific test statistic
	pValue := chiDist.Prob(chiSquare)

	// Output the p-value to standard out.
	fmt.Printf("p-value: %0.4f\n\n", pValue)
}
