package main

import (
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

	table1, err := csvutil.Create("pop.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer table1.Close()

	table1.Writer.Comma = ';'
	table1.Writer.UseCRLF = false

	rows, err := table.ReadRows(1, -1)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		data := State{}

		err = rows.Scan(&data)
		if err != nil {
			log.Println(err)
			continue
		}
		err = table1.WriteRow(data)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
	}
	table1.WriteHeader("UF;Pop 2010;Pop 2021")
}
