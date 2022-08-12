package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func printData(raw [][]string) {
	for i, line := range raw {
		str := fmt.Sprintf("%2d. [", i)
		for _, v := range line {
			str += fmt.Sprintf("%v ", v)
		}
		str = str[:len(str)-1] + "]"
		fmt.Println(str)
	}
}

func main() {
	os.Chdir("/home/rudson/Documents/Estudos/DeepLearning/csv/readCSV")

	// Open the irir dataset file
	f, err := os.Open("iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a ne CSV reader reading from the opened file
	reader := csv.NewReader(f)

	// Assume we don't know the number of fields per line. By setting
	// FieldsPerRecord negative, each row may have a variable
	// number of fields
	// reader.FieldsPerRecord = -1
	reader.FieldsPerRecord = 5

	// rawCSVData will hold our successfully parser rows
	var rawCSVData [][]string

	// Read in the records one by one
	for {
		// read in a row. check if we are at the end of file
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// if we had a parsing error, log the error and move on
		if err != nil {
			log.Println(err)
			continue
		}

		// append the record to our dataset
		rawCSVData = append(rawCSVData, record)
		fmt.Println(record)
	}

	fmt.Println("\nraw CVS Data:")
	printData(rawCSVData)
}
