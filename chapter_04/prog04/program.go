package main

import (
	"bufio"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func main() {
	f, err := os.Open("Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	advertDF := dataframe.ReadCSV(f)

	// Calculate the number of elements in each set
	trainingNum := (4 * advertDF.Nrow()) / 5
	testNum := advertDF.Nrow() - trainingNum

	// Create the subset indices
	trainingIdx := make([]int, trainingNum)
	testingIdx := make([]int, testNum)

	// Enumerate the training and testing indices
	for i := 0; i < trainingNum; i++ {
		trainingIdx[i] = i
	}
	for i := 0; i < testNum; i++ {
		testingIdx[i] = i + trainingNum
	}

	// Create the subset dataframes
	trainingDF := advertDF.Subset(trainingIdx)
	testDF := advertDF.Subset(testingIdx)

	// Create a map that will be used in writing the data to files
	setMap := map[int]dataframe.DataFrame{
		0: trainingDF,
		1: testDF,
	}

	// Create the respective files
	for idx, setName := range []string{"training.csv", "test.csv"} {
		// Save the filered dataset file
		f, err := os.Create(setName)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		// create a buffered writer
		w := bufio.NewWriter(f)

		// Write the dataframe out as a CSV
		if err := setMap[idx].WriteCSV(w); err != nil {
			log.Fatal(err)
		}
	}

}
