package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("labeled.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	var observed, predicted []int

	line := 1
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if line == 1 {
			line++
			continue
		}
		if err != nil {
			log.Println(err)
			log.Printf("Parsing line %d failed, unexpected type\n", line)
		}
		observedVal, err := strconv.Atoi(record[0])
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}
		predictedVal, err := strconv.Atoi(record[1])
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}
		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++
	}

	// Classes contains th tree possibile classes in the labeled data.
	classes := []int{0, 1, 2}
	truePos := make([]int, 3)
	falsePos := make([]int, 3)
	falseNeg := make([]int, 3)
	trueNeg := make([]int, 3)

	// Loop over esch class
	for idx, oVal := range observed {
		switch oVal {
		case 0:
			if predicted[idx] == oVal {
				truePos[0]++
				trueNeg[1]++
				trueNeg[2]++
			} else {
				falseNeg[0]++
				if predicted[idx] == 1 {
					falsePos[1]++
				} else {
					falsePos[2]++
				}
			}
		case 1:
			if predicted[idx] == oVal {
				truePos[1]++
				trueNeg[2]++
				trueNeg[0]++
			} else {
				falseNeg[1]++
				if predicted[idx] == 2 {
					falsePos[2]++
				} else {
					falsePos[0]++
				}
			}
		case 2:
			if predicted[idx] == oVal {
				truePos[2]++
				trueNeg[1]++
				trueNeg[0]++
			} else {
				falseNeg[2]++
				if predicted[idx] == 0 {
					falsePos[0]++
				} else {
					falsePos[1]++
				}
			}
		}
	}

	// Commum metrics:
	// The percentage of predictiions that were right:
	// (TP+TN)/(TP+TN+FP+FN)
	accuracy := make([]float64, 3)
	// The percentage of positive predictions that ware actually positive:
	// TP/(TP+FP)
	precision := make([]float64, 3)
	// The percentage og positive predictions that ware identified
	// as positive:
	// TP/(TP + FN)
	recall := make([]float64, 3)

	for _, class := range classes {
		accuracy[class] = float64(truePos[class]+trueNeg[class]) / float64(truePos[class]+trueNeg[class]+falsePos[class]+falseNeg[class])
		precision[class] = float64(truePos[class]) / float64(truePos[class]+falsePos[class])
		recall[class] = float64(truePos[class]) / float64(truePos[class]+falseNeg[class])

		fmt.Printf("\n----------------------------\n")
		fmt.Printf("Class: %d (%d samples)\n", class, trueNeg[class]+truePos[class]+falseNeg[class]+falsePos[class])
		fmt.Printf("TP: %2d  FP: %2d\n", truePos[class], falsePos[class])
		fmt.Printf("FN: %2d  TN: %2d\n", falseNeg[class], trueNeg[class])

		fmt.Printf("\nConfusion Matrix:\n")
		fmt.Printf("\t\tPredicted\n")
		fmt.Printf("\t\tFraud  Not Fraud\n")
		fmt.Printf("Observed  Fraud   %2d      %2d\n", truePos[class], falseNeg[class])
		fmt.Printf("      Not Fraud   %2d      %d\n\n", falsePos[class], trueNeg[class])

		fmt.Printf("Accuracy: %.2f   precision: %.2f  recall: %.2f\n", accuracy[class], precision[class], recall[class])
	}
}
