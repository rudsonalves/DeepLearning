package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	scoreMax = 830.0
	scoreMin = 640.0
)

func main() {
	fin, err := os.Open("loan_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	reader := csv.NewReader(fin)
	reader.FieldsPerRecord = 2
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Outpuf file
	fout, err := os.Create("clean_loan_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer fout.Close()

	w := csv.NewWriter(fout)

	// Sequentially move the rows writing out the parsed values
	for idx, record := range rawCSVData {
		// Header files
		if idx == 0 {
			// Write the header to the output file
			if err := w.Write(record); err != nil {
				log.Fatal(err)
			}
			continue
		}

		// Initialization a slice to hold our parser values
		outRecord := make([]string, 2)

		// Parse and standardize the FICO score
		score, err := strconv.ParseFloat(strings.Split(record[0], "-")[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		outRecord[0] = strconv.FormatFloat((score-scoreMin)/(scoreMax-scoreMin), 'f', 4, 64)

		// Parse the Interest rate class
		rate, err := strconv.ParseFloat(strings.TrimSuffix(record[1], "%"), 64)
		if err != nil {
			log.Fatal(err)
		}

		if rate <= 12.0 {
			outRecord[1] = "1.0"
		} else {
			outRecord[1] = "0.0"
		}

		// Write the record to the output file
		if err := w.Write(outRecord); err != nil {
			log.Fatal(err)
		}
	}

	// Write any buffered data to the underlying write
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
