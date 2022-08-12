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
	f, err := os.Open("iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	irisDF := dataframe.ReadCSV(f)

	// Create the plot from CSV file
	p := plot.New()

	// Create the plot and set its title and axis label
	p.Title.Text = "Box plots"
	p.Y.Label.Text = "Values"

	// Create the box for our data
	w := vg.Points(50)

	// Create a box plot for eachs of the features columns in the dataset
	for idx, colName := range irisDF.Names() {
		if colName != "species" {
			v := make(plotter.Values, irisDF.Nrow())
			for i, floatVal := range irisDF.Col(colName).Float() {
				v[i] = floatVal
			}

			b, err := plotter.NewBoxPlot(w, float64(idx), v)
			if err != nil {
				log.Fatal(err)
			}
			p.Add(b)
		}
	}
	p.NominalX("sepal_length", "sepal_width", "petal_length", "petal_width")
	if err := p.Save(6*vg.Inch, 8*vg.Inch, "boxplots.png"); err != nil {
		log.Fatal(err)
	}
}
