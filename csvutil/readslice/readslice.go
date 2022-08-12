package main

import (
	"fmt"
	"log"

	"go-hep.org/x/hep/csvutil"
)

type State struct {
	UF      string
	Pop2010 int
	Pop2021 int
}

func main() {
	fcsv, err := csvutil.Open("populacao2.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer fcsv.Close()

	fcsv.Reader.Comma = ','
	fcsv.Reader.Comment = '#'

	rows, err := fcsv.ReadRows(0, -1)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	records := []State{}

	for rows.Next() {
		var (
			UF      string
			Pop2010 int
			Pop2021 int
		)
		err = rows.Scan(&UF, &Pop2010, &Pop2021)
		if err != nil {
			log.Println(err)
		}
		data := State{UF, Pop2010, Pop2021}
		records = append(records, data)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(records)
}
