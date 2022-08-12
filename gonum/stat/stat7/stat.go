package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat"
)

/*

A survey of local residents reveals that 60% of all residents get no regular exercise, 25% exercise sporadically, and 15% exercise regularly. After doing some fancy modeling and putting some community services in place, the survey was repeated with the same questions. The follow-up survey was completed by 500 residents with the following results:

- No regular exercise: 260
- Sporadic exercise: 135
- Regular exercise: 105
Total: 500

Now, we want to determine if there is evidence for a statistically significant shift in the responses of the residents. Our null and alternate hypotheses are as follows:

- H0: The deviations from the previously observed percentages are due to pure
chance
- Ha: The deviations are due to some underlying effect outside of pure chance
(possibly our new community services)

*/

func main() {
	observed := []float64{260, 135, 105}
	totalObserved := func(values []float64) float64 {
		s := 0.
		for _, v := range values {
			s += v
		}
		return s
	}(observed)

	expected := []float64{
		totalObserved * .60,
		totalObserved * .25,
		totalObserved * .15,
	}

	chiSquare := stat.ChiSquare(observed, expected)
	fmt.Printf("\nChi-square: %0.2f\n", chiSquare)
}
