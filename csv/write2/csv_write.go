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
	defer w.Flush()

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Print(err)
		}
	}
}
