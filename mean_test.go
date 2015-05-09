package calc

import "testing"

func TestMean(t *testing.T) {
	exp := 4.85
	act := Mean([]float64{2.3, 3.5, 5.7, 7.9})
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
