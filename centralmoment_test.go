package calc

import (
	"fmt"
	"testing"
)

func TestCentralMoment(t *testing.T) {
	vals := []float64{2.3, 3.5, 5.7, 7.9}
	exps := []float64{1.33333333, 0.0, 6.11666667, 3.315, 44.22084167}

	// Test each central moment from 0 to 4.
	for i := 0; i < 5; i++ {
		act := CentralMoment(vals, i)
		if !Near(exps[i], act) {
			fmt.Print(exps[i] - act)
			t.Error("Expected", exps[i], "got", act)
		}
	}
}
