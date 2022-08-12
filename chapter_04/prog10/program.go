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

	// Lopp over the test data predicting y and evaluating the
	// prediction with the mean absolute error
	var mAE float64
	var nTestData = float64(len(testData))
	for i, record := range testData {
		// Skip the header
		if i == 0 {
			continue
		}

		// Parse Sales
		yObserved, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}
		// Parse TV
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}
		// Parse Radio
		radioVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}
		// Parse NewsPaper
		newsPaper, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Predict y with our trained model
		yPredited := Predicted(tvVal, radioVal, newsPaper)

		// Add the to the mean absolute error
		mAE += math.Abs(yObserved-yPredited) / nTestData
	}

	// Output the MAE to standard out
	fmt.Printf("MAE = %0.2f\n\n", mAE)
}

func Predicted(tv, radio, newspaper float64) float64 {
	return 2.9490 + tv*0.0473 + radio*0.1799 + newspaper*-0.0009
}
