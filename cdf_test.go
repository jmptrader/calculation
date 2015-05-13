package calc

import "testing"

func TestCDF(t *testing.T) {
	exps := []float64{0.97724994, 0.84134474, 0.5, 0.84134474, 0.97724994}
	vals := []float64{-2.0, -1.0, 0.0, 1.0, 2.0}
	for i, exp := range exps {
		act := CDF(vals[i])
		if !Near(exp, act) {
			t.Error("Expected", exp, "got", act)
		}
	}
}
