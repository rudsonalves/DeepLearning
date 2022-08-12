package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-gota/gota/series"
	"github.com/kniren/gota/dataframe"
)

func main() {
	f, err := os.Open("populacao.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)
	if err := df.Error(); err != nil {
		fmt.Println(err)
	}

	// ------ Estados com menos de 1M de habitantes -------
	fmt.Println("Estados com menos de 1 milhão de habitantes:")
	f1 := dataframe.F{
		Colname:    "População (2021)",
		Comparator: "<",
		Comparando: 1000000,
	}

	df1 := df.Filter(f1)
	if df1.Err != nil {
		fmt.Println(err)
	}
	fmt.Println(df1)

	// ------ Estados iniciados com R -------
	fmt.Println("Estados iniciados com a letra \"R\":")
	fsearch := func(prefix string) func(series.Element) bool {
		return func(e series.Element) bool {
			if e.Type() == series.String {
				if str, ok := e.Val().(string); ok {
					return strings.HasPrefix(str, prefix)
				}
			}
			return false
		}
	}
	f2 := dataframe.F{
		Colname:    "Unidade Federativa",
		Comparator: series.CompFunc,
		Comparando: fsearch("R"),
	}

	df2 := df.Filter(f2)
	if df2.Err != nil {
		fmt.Println(err)
	}
	fmt.Println(df2)

	// ------ Estados com menos de 1M de habitantes e iniciados por R -------
	fmt.Println("Estados iniciados com a letra \"R\" e com menos de 1 milhão de habitantes:")
	df3 := df1.Filter(f2)
	if df3.Err != nil {
		fmt.Println(err)
	}
	fmt.Println(df3)

	df4 := df.FilterAggregation(dataframe.And, f1, f2)
	fmt.Println(df4)

	fmt.Println(df.Arrange(dataframe.Sort("Unidade Federativa")))
	fmt.Println(df.Arrange(dataframe.RevSort("População (2021)")))
}
