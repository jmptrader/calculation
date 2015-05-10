package calc

import "testing"

func TestTrimmedMean(t *testing.T) {
	exp := 6.25555556
	act := TrimmedMean([]float64{5.2, 9.1, 9.8, 9.5, 0.4, 6.7, 5.7, 3.2, 6.7}, 0.1)
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
