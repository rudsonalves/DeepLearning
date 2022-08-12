package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/sajari/regression"
)

func main() {
	f, err := os.Open("training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 4
	trainingData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	/* In this code we will create a multiple linear regression
	model that looks like:

	Sales = m1 TV + m2 Radio + b

	*/
	var r regression.Regression
	r.SetObserved("Sales")
	r.SetVar(0, "TV")
	r.SetVar(1, "Radio")
	r.SetVar(2, "Newspaper")

	// Loop over the CSV records adding the training data
	for i, record := range trainingData {
		// skip the header
		if i == 0 {
			continue
		}

		// Parse the Sales
		yVal, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}
		// Parse TV value
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}
		// Parse Radio value
		radioVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}
		// Parse Newspaper
		newsPaper, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Add these points to the regression value
		r.Train(regression.DataPoint(yVal, []float64{tvVal, radioVal, newsPaper}))
	}

	// Train/fit the regression model
	r.Run()

	// Output the trained model parameters
	fmt.Printf("\nRegresssion Formula: \n%v\n\n", r.Formula)
}
