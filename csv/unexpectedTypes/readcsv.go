package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// CSVRecord contains a sucessfully parser row of the CSV file
type CSVRecord struct {
	SepalLength float64
	SepalWidth  float64
	PetalLength float64
	PentalWidth float64
	Species     string
	ParseError  error
}

func printData(raw []CSVRecord) {
	for i, l := range raw {
		str := fmt.Sprintf("%2d. [%.1f %.1f %.1f %.1f %q]", i, l.SepalLength, l.SepalWidth, l.PetalLength, l.PentalWidth, l.Species)
		fmt.Println(str)
	}
}

func main() {
	os.Chdir("/home/rudson/Documents/Estudos/DeepLearning/csv/unexpectedTypes")

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

	// Create a slice value that will hold all of the sucessfully parsed
	// records from the CSV
	var csvData []CSVRecord

	// Read in the records one by one
	for {
		// read in a row. check if we are at the end of file
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// Create a CVSRecord value for the row
		var csvRecord CSVRecord

		for id, value := range record {
			// Parse the value in the record as a string for the string column
			if id == 4 {
				// Validate that the value is not an empty string. If the
				// value is an empty string break the parsing loop.
				if value == "" {
					log.Printf("Unexpected type in the column %d\n", id)
					csvRecord.ParseError = fmt.Errorf("empty string value")
					break
				}

				// Add the string value to the CVSRecord
				csvRecord.Species = value
				continue
			}

			// Otherwise, parse the value in the record as a float64
			var floatValue float64

			// If the value can not be parsed as a float64, log and break the
			// parsing loop
			floatValue, err = strconv.ParseFloat(value, 64)
			if err != nil {
				log.Printf("Unexpected type in column %d\n", id)
				csvRecord.ParseError = fmt.Errorf("could not parse float")
				break
			}

			// Add the float value to the respective field in the CSVRecord
			switch id {
			case 0:
				csvRecord.SepalLength = floatValue
			case 1:
				csvRecord.SepalWidth = floatValue
			case 2:
				csvRecord.PetalLength = floatValue
			case 3:
				csvRecord.PentalWidth = floatValue
			}
		}
		// Append successfully parsed records to the slice defined above
		if csvRecord.ParseError == nil {
			csvData = append(csvData, csvRecord)
		}
	}

	fmt.Println("\nraw CVS Data:")
	printData(csvData)
}
