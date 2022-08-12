package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type State struct {
	name    string
	pop2010 int
	pop2021 int
}

func printReport(States []State) {
	total2010, total2021 := 0, 0
	for _, state := range States {
		total2010 += state.pop2010
		total2021 += state.pop2021
	}
	fmt.Printf("   %20s %9s  %9s  %5s  %5s\n", "Estado", "Pop 2010", "Pop 2021", "%2010", "%2021")
	for i, state := range States {
		pp2010 := 100. * float64(state.pop2010) / float64(total2010)
		pp2021 := 100. * float64(state.pop2021) / float64(total2021)
		fmt.Printf("%2d %20s %9d  %9d  %4.1f%%  %4.1f%%\n", i+1, state.name, state.pop2010, state.pop2021, pp2010, pp2021)
	}
}

func main() {
	f, err := os.Open("populacao.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var brasil []State

	records := csv.NewReader(f)
	records.FieldsPerRecord = 3

	for {
		record, err := records.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Print(err)
			continue
		}
		nome := record[0]
		if record[1] == "População (2010)" {
			continue
		}
		pop10, err := strconv.Atoi(record[1])
		if err != nil {
			log.Print(err)
			continue
		}
		pop21, err := strconv.Atoi(record[2])
		if err != nil {
			log.Print(err)
			continue
		}
		brasil = append(brasil, State{name: nome, pop2010: pop10, pop2021: pop21})
	}

	printReport(brasil)
}
