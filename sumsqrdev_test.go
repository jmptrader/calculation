package calc

import "testing"

func TestSumSqrDev(t *testing.T) {
	exp := 18.35
	act := SumSqrDev([]float64{2.3, 3.5, 5.7, 7.9})
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
