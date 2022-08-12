package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/montanaflynn/stats"
	"gonum.org/v1/gonum/stat"
)

func main() {
	irisFile, err := os.Open("iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	irisDF := dataframe.ReadCSV(irisFile)

	sepalLength := irisDF.Col("petal_length").Float()

	meanVal := stat.Mean(sepalLength, nil)

	modeVal, modeCount := stat.Mode(sepalLength, nil)

	medianVal, err := stats.Median(sepalLength)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nSepal Length Summary Statistics:\n")
	fmt.Printf("Mean value: %0.2f\n", meanVal)
	fmt.Printf("Mode value: %0.2f\n", modeVal)
	fmt.Printf("Mode count: %d\n", int(modeCount))
	fmt.Printf("Median value: %0.2f\n\n", medianVal)
	// fmt.Println(irisDF)
}
