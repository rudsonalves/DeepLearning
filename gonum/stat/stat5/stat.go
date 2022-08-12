package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/montanaflynn/stats"
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
	sepalLength := irisDF.Col("sepal_length").Float()
	sepalWidth := irisDF.Col("sepal_width").Float()
	cor, _ := stats.Correlation(sepalLength, sepalWidth)
	fmt.Println(cor)

	pts := make(plotter.XYs, len(sepalLength))

	for i, floatVal := range sepalLength {
		pts[i].X = floatVal
		pts[i].Y = sepalWidth[i]
	}

	p := plot.New()
	p.X.Label.Text = "sepalLength"
	p.Y.Label.Text = "sepalWidth"
	p.Add(plotter.NewGrid())
	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.GlyphStyle.Radius = vg.Points(3)
	p.Add(s)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "scatter.png"); err != nil {
		log.Fatal(err)
	}
}
