package main

import (
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func predict(tv float64) float64 {
	return 7.0688 + tv*.04899
}

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

	// pts will hold the values for plotting
	pts := make(plotter.XYs, advertDF.Nrow())

	// ptsPred will hold the predicted values for plotting
	ptsPred := make(plotter.XYs, advertDF.Nrow())

	// Fill pts with data
	for i, floatVal := range advertDF.Col("TV").Float() {
		pts[i].X = floatVal
		pts[i].Y = yVals[i]
		ptsPred[i].X = floatVal
		ptsPred[i].Y = predict(floatVal)
	}

	// Create the plot
	p := plot.New()
	p.X.Label.Text = "TV"
	p.Y.Label.Text = "Sales"
	p.Add(plotter.NewGrid())

	// Add the scatter plot points for the observations
	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}
	s.GlyphStyle.Radius = vg.Points(3)

	// Add the line plot points for the predictions
	l, err := plotter.NewLine(ptsPred)
	if err != nil {
		log.Fatal(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}

	//  Save the plot to a PNG file
	p.Add(s, l)
	if err := p.Save(width, length, "regression_line.png"); err != nil {
		log.Fatal(err)
	}
}
