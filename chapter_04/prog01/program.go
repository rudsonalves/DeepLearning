package main

import (
	"fmt"
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

	// Use the Descripe method to calculate summary statistics
	// for all of the columns in one shot
	advertSummary := advertDF.Describe()

	// Output the summary statistics to stdout.
	fmt.Println(advertSummary)
}
