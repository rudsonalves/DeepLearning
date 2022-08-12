package main

import (
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

	// Extract the target column
	yVals := advertDF.Col("Sales").Float()

	// Create a scatter plot for each of the features in the dataset
	for _, colName := range advertDF.Names() {
		// pts will hold the values for plotting
		pts := make(plotter.XYs, advertDF.Nrow())

		// Fill pts with data
		for i, floatVal := range advertDF.Col(colName).Float() {
			pts[i].X = floatVal
			pts[i].Y = yVals[i]
		}

		// Create the plot
		p := plot.New()
		p.X.Label.Text = colName
		p.Y.Label.Text = "y"
		p.Add(plotter.NewGrid())

		s, err := plotter.NewScatter(pts)
		if err != nil {
			log.Fatal(err)
		}
		s.GlyphStyle.Radius = vg.Points(3)

		// Save the plot to a PNG file
		p.Add(s)
		if err := p.Save(width, length, colName+"_scatter.png"); err != nil {
			log.Fatal(err)
		}
	}
}
