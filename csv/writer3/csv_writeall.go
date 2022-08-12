package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	var records = [][]string{
		{"first_name", "last_name", "age"},
		{"Rob", "Pike", "31"},
		{"Ken", "Thompson", "46"},
		{"Robert", "Griesemer", "28"},
		{"Albert", "Eintein", "143"}}

	f, err := os.Create("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.Comma = ';'
	w.UseCRLF = false

	w.WriteAll(records)
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
