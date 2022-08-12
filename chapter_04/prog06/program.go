package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 4
	testData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Loop over the test data predicting y and evaluating the prediction
	// with the mean absolute error.
	var mAE float64
	for i, record := range testData {
		// Skip the header
		if i == 0 {
			continue
		}

		// Parse the observed Sales, or "y"
		yObserved, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Parse the TV value
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Predict y with our trained model
		yPredicted := predict(tvVal)

		// Add the to the mean absolute error.
		mAE += math.Abs(yObserved-yPredicted) / float64(len(testData))
	}
	// Output the MAE to standard out.
	fmt.Printf("MAE = %0.2f\n\n", mAE)
}

func predict(tv float64) float64 {
	return 7.0688 + tv*.04899
}
