package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/gonum/floats"
)

type centroid []float64

// dfFloatRow retrieves a slice of float values from a DataFrame
// at the given index and for the given column names
func dfFloatRow(df dataframe.DataFrame, names []string, idx int) []float64 {
	var row []float64
	for _, name := range names {
		row = append(row, df.Col(name).Float()[idx])
	}
	return row
}

func main() {
	f, err := os.Open("iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	irisDF := dataframe.ReadCSV(f)

	// Define the names of the three separate species cotaine in the CSV
	speciesNames := []string{
		"Iris-setosa",
		"Iris-versicolor",
		"Iris-virginica",
	}

	// Create a map to hold our cntroid information
	centroids := make(map[string]centroid)

	// Create a map to hold the filtered dataframe for each cluster
	clusters := make(map[string]dataframe.DataFrame)

	// Filter the dataset into three separate dataframes,
	// each corresponing to one o the Irir species
	for _, species := range speciesNames {
		// filter the original dataset
		filter := dataframe.F{
			Colname:    "species",
			Comparator: "==",
			Comparando: species,
		}
		filtered := irisDF.Filter(filter)

		// Add the filtered dataframe to the map of clusters
		clusters[species] = filtered

		// Calculate the mean of features
		summaryDF := filtered.Describe()

		// fmt.Println(species)
		// fmt.Println(summaryDF.Names())

		// Put each dimension's mean into the corresponding centroid
		var c centroid
		for _, feature := range summaryDF.Names() {
			// Skip the irrelevant columns
			if feature == "column" || feature == "species" {
				continue
			}
			c = append(c, summaryDF.Col(feature).Float()[0])
		}

		// Add this centroid to our map
		centroids[species] = c
	}

	// Convert our labels into a slice of strings and create a slice
	// of float column names for convenience
	labels := irisDF.Col("species").Records()
	floatColumns := []string{
		"sepal_length",
		"sepal_width",
		"petal_length",
		"petal_width",
	}

	// fmt.Println(labels)
	// fmt.Println(floatColumns)

	// Loop over the records accumulating the average silhouette coecifient
	var silhouette float64

	for idx, label := range labels {
		// a will store our accumulated value for a
		var a float64

		clusters_nrows := clusters[label].Nrow()

		// Get the current data point for comparison
		current := dfFloatRow(irisDF, floatColumns, idx)
		// Loop over the data points in the came cluster
		for i := 0; i < clusters_nrows; i++ {
			// Get the data point for comparison
			other := dfFloatRow(clusters[label], floatColumns, i)

			// add to a
			a += floats.Distance(current, other, 2) / float64(clusters_nrows)
		}
		// Dtermine the nearest other cluster
		var otherCluster string
		var distanceToCluster float64
		for _, species := range speciesNames {
			// Skip the cluster containing the data point
			if species == label {
				continue
			}

			// Calculate the distance to the cluster from the curret cluster
			distanceFrothisCluster := floats.Distance(centroids[label], centroids[species], 2)

			// Replace the current cluster if relevant
			if distanceToCluster == 0.0 || distanceFrothisCluster < distanceToCluster {
				otherCluster = species
				distanceToCluster = distanceFrothisCluster
			}
		}

		// b will store our accumulated for b
		var b float64
		// Get the data point for comparison
		// current := dfFloatRow(irisDF, floatColumns, idx)
		for i := 0; i < clusters[otherCluster].Nrow(); i++ {
			// Get the data point for comparison
			other := dfFloatRow(clusters[otherCluster], floatColumns, i)

			// Add to b
			b += floats.Distance(current, other, 2) / float64(clusters[otherCluster].Nrow())
		}

		// Add to the average silhouette coefficient
		if a > b {
			silhouette += ((b - a) / a) / float64(len(labels))
		}
		silhouette += ((b - a) / b) / float64(len(labels))
	}

	// as a sanity check, output our centroids
	// for _, species := range speciesNames {
	// 	fmt.Printf("%s centroid: %v\n", species, centroids[species])
	// }
	// Output the final average silhouette coefficient to stdout.
	fmt.Printf("\nAverage Silhouette Coefficient: %0.2f\n\n", silhouette)
}
