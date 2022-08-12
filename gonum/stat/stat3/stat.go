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
	f, err := os.Open("iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	irisDF := dataframe.ReadCSV(f)

	// Create a histogram for each of the feature columns in the dataset
	for _, colName := range irisDF.Names() {
		// if the columns is one of the features columns, let's create
		// a histogram ot eht values
		if colName != "species" {
			// Take column values
			v := make(plotter.Values, irisDF.Nrow())
			for i, floatVal := range irisDF.Col(colName).Float() {
				v[i] = floatVal
			}

			// Make a plot and set its title
			p := plot.New()
			p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)

			// Create a histogram of our values drawn
			// from the standard normal
			h, err := plotter.NewHist(v, 16)
			if err != nil {
				log.Fatal(err)
			}

			// Normalize the histogram
			h.Normalize(1)

			// add the histogram to the plot
			p.Add(h)

			// Save the plot to a PNG file
			if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"); err != nil {
				log.Fatal(err)
			}
		}
	}
}
