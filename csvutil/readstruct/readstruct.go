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
	table, err := csvutil.Open("populacao.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer table.Close()

	table.Reader.Comma = ','
	table.Reader.Comment = '#'

	rows, err := table.ReadRows(1, -1)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	records := []State{}

	for rows.Next() {
		data := State{}

		err = rows.Scan(&data)
		if err != nil {
			log.Println(err)
			continue
		}
		records = append(records, data)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(records)
}
