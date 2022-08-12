package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func main() {
	os.Chdir("/home/rudson/Documents/Estudos/DeepLearning/csv/dataframe")

	// Open CSV file
	irisFile, err := os.Open("iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Create a dataframe from the CSV file
	// The types of the columns will be inferred
	irisDF := dataframe.ReadCSV(irisFile)

	// As a sanity check, display the recosds to stdout
	// Gota will format the dataframe for pretty printing
	fmt.Println(irisDF)

	// Create a filter for the dataframe
	filter := dataframe.F{
		Colname:    "species",
		Comparator: "==",
		Comparando: "Iris-versicolor",
	}

	// Filter the dataframe to see only the rows where
	// the irir species is "Iris-"-versicolor
	versicolorDF := irisDF.Filter(filter)
	if versicolorDF.Err != nil {
		log.Fatal(versicolorDF.Err)
	}
	fmt.Println(versicolorDF)

	// Filter the dataframe again, but only select out the
	// sepal_width and species columns
	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_width", "species"})
	fmt.Println(versicolorDF)

	// Filter and select the dataframe again, but only display
	// the first tree results
	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_width", "species"}).Subset([]int{0, 1})
	fmt.Println(versicolorDF)
}
