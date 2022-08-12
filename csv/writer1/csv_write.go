package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

const input = `first_name,last_name,age
Rob,Pike,31
Ken,Thompson,46
Robert,Griesemer,28
Albert,Eintein,143
`

func main() {
	r := csv.NewReader(strings.NewReader(input))
	allRec, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.Comma = ','
	w.UseCRLF = false
	defer w.Flush()

	for _, record := range allRec {
		if err := w.Write(record); err != nil {
			log.Print(err)
		}
	}
}
