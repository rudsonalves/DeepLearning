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
	f, err := os.Open("clean_loan_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	loanDF := dataframe.ReadCSV(f)

	// Use the Describe method to calculate summary statistics
	loanSummary := loanDF.Describe()
	fmt.Println(loanSummary)

	// Create a histogram for each of the columns in the dataset
	for _, colName := range loanDF.Names() {
		// Create a plotter.Values value and fill it with the values
		plotVals := make(plotter.Values, loanDF.Nrow())
		for i, floatVal := range loanDF.Col(colName).Float() {
			plotVals[i] = floatVal
		}

		// Make a plot and set its title
		p := plot.New()
		p.Title.Text = fmt.Sprintf("Histogram og a %s", colName)

		// Create a histogram of out values
		h, err := plotter.NewHist(plotVals, 16)
		if err != nil {
			log.Fatal(err)
		}
		h.Normalize(1)

		// Add the histogram to the plot
		p.Add(h)

		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"); err != nil {
			log.Fatal(err)
		}
	}
}
