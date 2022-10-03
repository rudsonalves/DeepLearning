package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

/*

Testing program

*/

// predict makes a prediction based on our
// trained logistic regression model
func predict(score float64) float64 {
	p := 1 / (1 + math.Exp(-13.64*score+4.88))

	// Output the corresponding class
	if p >= 0.5 {
		return 1.0
	}
	return 0.0
}

func main() {
	f, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	var observed []float64
	var predicted []float64

	// line will track row numbers for logging
	line := 1

	// Read in the records lookin for unexpected types in the columns
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		// Skil the header
		if line == 1 {
			line++
			continue
		}
		// Read in the observed value
		observedVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Make the corresponding prediction
		score, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		predictedVal := predict(score)

		// Append the record to our slice, if it has the expected type
		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++
	}

	// This variable will hold our count of true positive and
	// true negative values
	var truePosNeg int

	//  Accumulate the true positive/negative count
	for idx, oVal := range observed {
		if oVal == predicted[idx] {
			truePosNeg++
		}
	}

	// Calculate the accuracy (subset accuracy)
	accuracy := float64(truePosNeg) / float64(len(observed))

	// Output the Accuracy value to standard out
	fmt.Printf("\nAccuracy = %0.2f\n\n", accuracy)
}
