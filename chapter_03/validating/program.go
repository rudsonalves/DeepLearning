package main

import (
	"bufio"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func main() {
	f, err := os.Open("diabetes.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	diabetesDF := dataframe.ReadCSV(f)

	// Calculate the number of elements in each set
	trainingNum := (4 * diabetesDF.Nrow()) / 5
	testeNum := diabetesDF.Nrow() / 5
	if trainingNum+testeNum < diabetesDF.Nrow() {
		trainingNum++
	}

	// Create the subset indices
	trainingIdx := make([]int, trainingNum)
	testIdx := make([]int, testeNum)

	// Enumerate the training indices
	for i := 0; i < trainingNum; i++ {
		trainingIdx[i] = i
	}

	// Enumerate the test indices
	for i := 0; i < testeNum; i++ {
		testIdx[i] = i + trainingNum
	}

	// Create the subset dataframes
	traininigDF := diabetesDF.Subset(trainingIdx)
	testDF := diabetesDF.Subset(testIdx)

	// Create a map that will be used in writing the data
	// to files
	setMap := map[int]dataframe.DataFrame{
		0: traininigDF,
		1: testDF,
	}

	// Create the respective files
	for idx, setName := range []string{"training.csv", "test.csv"} {
		f, err := os.Create(setName)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		// Create a buffered writer
		w := bufio.NewWriter(f)

		// Write the dataframe out as a CSV.
		if err := setMap[idx].WriteCSV(w); err != nil {
			log.Fatal(err)
		}
	}
}
