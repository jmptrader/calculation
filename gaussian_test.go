package calc

import "testing"

func TestGaussian(t *testing.T) {
	exps := []float64{0.0044318484, 0.05399097, 0.24197072, 0.39894228, 0.24197072}
	vals := []float64{-2.0, -1.0, 0.0, 1.0, 2.0}
	for i, exp := range exps {
		act := Gaussian(vals[i], 1, 1)
		if !Near(exp, act) {
			t.Error("Expected", exp, "got", act)
		}
	}
}
