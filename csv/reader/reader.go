package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

func strSlice(list []string) string {
	str := "["
	for _, s := range list {
		str += fmt.Sprintf("%q ", s)
	}
	str = str[:len(str)-1] + "]"
	return str
}

const input = `first_name,last_name,username
"Rob","Pike",'rob'
Ken,Thom"pson,ken
"Robert","Griesemer","gri"
# lines beginning with a # character are ignored if r.Comment = #
   Campo 1 ,   Campo 2   ,    Campo 3
Campo 1; Campo 2; Campo 3
I Don't Know, I Don't Know, I Don't Know
`

func main() {
	r := csv.NewReader(strings.NewReader(input))
	r.Comma = ','
	r.Comment = '#'
	r.FieldsPerRecord = 3
	r.LazyQuotes = true
	r.TrimLeadingSpace = true
	r.ReuseRecord = true

	// var records [][]string
	// for {
	// 	record, err := r.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Println(err)
	// 		continue
	// 	}
	// 	records = append(records, record)
	// }

	// for i, line := range records {
	// 	s := strSlice(line)
	// 	fmt.Printf("%2d: %s\n", i, s)
	// }

	var i int = 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			continue
		}
		s := strSlice(record)
		fmt.Printf("%2d: %s\n", i, s)
		i++
	}
}
