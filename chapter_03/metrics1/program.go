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

	// This variable will hold out count ot true positive and true negative values
	var truePosNeg int

	// Accumulate the true positive/negative count
	for idx, oVal := range observed {
		if oVal == predicted[idx] {
			truePosNeg++
		}
	}

	// Calculate the accuracy (subset accuracy)
	accuracy := float64(truePosNeg) / float64(len(observed))

	// Output the Accuracy value to standard out.
	fmt.Printf("\nAccuracy = %0.2f\n\n", accuracy)

	// Classes contains th tree possibile classes in the labeled data.
	classes := []int{0, 1, 2}

	// Loop over esch class
	for _, class := range classes {
		var truePos, falsePos, falseNeg int

		// Accumulate the true positive is the relevant positive counts
		for idx, oVal := range observed {

			switch oVal {
			// if the observed value is the relevant class, we should
			// check to see if we predicted that class
			case class:
				if predicted[idx] == class {
					truePos++
					continue
				}
				falseNeg++
			// if the observed value is a different class, we should
			// check to see if we predicted a false positive
			default:
				if predicted[idx] == class {
					falsePos++
				}
			}
		}
		// Calculate the precision.
		precision := float64(truePos) / float64(truePos+falsePos)

		// Calculate the recall
		recall := float64(truePos) / float64(truePos+falseNeg)

		// Output the precision value to standard out.
		fmt.Printf("\nPrecision (class %d) = %0.2f", class, precision)
		fmt.Printf("\nRecall (class %d) = %0.2f\n\n", class, recall)
	}

	for _, class := range []int{0, 1, 2} {
		var truePos, falsePos, falseNeg int
		for idx, oVal := range observed {
			switch oVal {
			case class:
				if predicted[idx] == class {
					truePos++
					continue
				} else {
					falseNeg++
				}
			default:
				if predicted[idx] == class {
					falsePos++
				}
			}
		}
		fmt.Printf("Class: %d\n", class)
		fmt.Printf("truePos:  %2d  falsePos: %2d\n", truePos, falsePos)
		fmt.Printf("falseNeg: %2d\n\n", falseNeg)
	}

	var TP0, TP1, TP2, FP0, FP1, FP2, FN0, FN1, FN2 int
	for idx, oVal := range observed {
		switch oVal {
		case 0:
			if predicted[idx] == oVal {
				TP0++
			} else {
				FN0++
				if predicted[idx] == 1 {
					FP1++
				} else {
					FP2++
				}
			}
		case 1:
			if predicted[idx] == oVal {
				TP1++
			} else {
				FN1++
				if predicted[idx] == 0 {
					FP0++
				} else {
					FP2++
				}
			}
		case 2:
			if predicted[idx] == oVal {
				TP2++
			} else {
				FN2++
				if predicted[idx] == 0 {
					FP0++
				} else {
					FP1++
				}
			}
		}
	}

	var TN0, TN1, TN2 int
	TN0 = TP1 + TP2
	TN1 = TP0 + TP2
	TN2 = TP0 + TP1
	fmt.Printf("Confusion Matrix to Class %d:\n", 0)
	fmt.Printf("\t\tPredicted\n")
	fmt.Printf("\t\tFraud  Not Fraud\n")
	fmt.Printf("Observed  Fraud   %2d      %2d\n", TP0, FN0)
	fmt.Printf("      Not Fraud   %2d      %d\n\n", FP0, TN0)

	fmt.Printf("Confusion Matrix to Class %d:\n", 1)
	fmt.Printf("\t\tPredicted\n")
	fmt.Printf("\t\tFraud  Not Fraud\n")
	fmt.Printf("Observed  Fraud   %2d      %2d\n", TP1, FN1)
	fmt.Printf("      Not Fraud   %2d      %d\n\n", FP1, TN1)

	fmt.Printf("Confusion Matrix to Class %d:\n", 2)
	fmt.Printf("\t\tPredicted\n")
	fmt.Printf("\t\tFraud  Not Fraud\n")
	fmt.Printf("Observed  Fraud   %2d      %2d\n", TP2, FN2)
	fmt.Printf("      Not Fraud   %2d      %d\n\n", FP2, TN2)
}
