package calc

import "testing"

func TestWeightedMean(t *testing.T) {
	exp := 4.425
	act := WeightedMean([]float64{2.3, 3.5, 5.7, 7.9}, []float64{3.0, 2.0, 1.0, 2.0})
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
