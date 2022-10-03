package main

import (
	"bufio"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

/*
Creating a training and test sets
*/
const inputCSVFile = "clean_loan_data.csv"

var outputCSVFiles = []string{"training.csv", "test.csv"}

func main() {
	f, err := os.Open(inputCSVFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	loanDF := dataframe.ReadCSV(f)

	// Number of elements in each set
	trainingNum := (4 * loanDF.Nrow()) / 5
	testNum := loanDF.Nrow() - trainingNum

	// Create the subset indices
	trainingIdx := make([]int, trainingNum)
	testIdx := make([]int, testNum)

	for i := 0; i < trainingNum; i++ {
		trainingIdx[i] = i
	}
	for i := 0; i < testNum; i++ {
		testIdx[i] = trainingNum + i
	}

	// Create subset dataframes
	trainingDF := loanDF.Subset(trainingIdx)
	testDF := loanDF.Subset(testIdx)

	// Create a map that will de used in writing the dta do files
	setMap := map[int]dataframe.DataFrame{
		0: trainingDF,
		1: testDF,
	}

	// Create the respective files
	for idx, setName := range outputCSVFiles {
		f, err := os.Create(setName)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		w := bufio.NewWriter(f)
		if err := setMap[idx].WriteCSV(w); err != nil {
			log.Fatal(err)
		}
	}
}
