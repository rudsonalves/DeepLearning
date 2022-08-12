package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"

	"gonum.org/v1/gonum/stat"
)

func main() {
	f, err := os.Open("continuous_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	var observed []float64
	var predicted []float64

	line := 1
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if line == 1 {
			line++
			continue
		}
		if err != nil {
			log.Println(err)
			continue
		}

		observedVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}
		predictedVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++
	}

	// Calculate the mean absolute error and mean squared error
	maxIdx := float64(len(observed))
	var mAE, mSE float64
	for idx, oVal := range observed {
		mAE += math.Abs(oVal-predicted[idx]) / maxIdx
		mSE += math.Pow(oVal-predicted[idx], 2) / maxIdx
	}

	// Output the MAE and MSE value to standard out.
	fmt.Printf("\nMAE = %0.2f\n", mAE)
	fmt.Printf("\nMSE = %0.2f\n", mSE)
	fmt.Printf("\nSMSE = %0.2f\n", math.Sqrt(mSE))

	// Calculate the R² value
	rSquared := stat.RSquaredFrom(observed, predicted, nil)

	// Output the R² value to standard out.
	fmt.Printf("\nR² = %0.2f\n\n", rSquared)
}
