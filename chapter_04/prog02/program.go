package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	const (
		width  = 4 * vg.Inch
		length = 4 * vg.Inch
	)

	f, err := os.Open("Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	advertDF := dataframe.ReadCSV(f)

	// Create a histogram for each of the columns in the dataset
	for _, colName := range advertDF.Names() {
		// Create a plotter.Values value and fill it with the
		// values from the respective column of the dataframe
		plotVals := make(plotter.Values, advertDF.Nrow())
		for i, floatVal := range advertDF.Col(colName).Float() {
			plotVals[i] = floatVal
		}

		// Make a plot and set its title
		p := plot.New()
		p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)

		// Create a histogram of our values drawn
		// from the standard nromal
		h, err := plotter.NewHist(plotVals, 16)
		if err != nil {
			log.Fatal(err)
		}

		// Normalize the histogram
		h.Normalize(1)

		// add the histogram to the plot
		p.Add(h)

		// Save the plot to a PNG file
		if err := p.Save(width, length, colName+"_hist.png"); err != nil {
			log.Fatal(err)
		}
	}
}
