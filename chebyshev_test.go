package calc

import "testing"

func TestChebyshev(t *testing.T) {
	exps := []float64{0.0, 0.0, 0.75, 0.88888889, 0.9375, 0.96}
	for i := 0; i < 6; i++ {
		act := Chebyshev(float64(i))
		if !Near(exps[i], act) {
			t.Error("Expected", exps[i], "got", act)
		}
	}
}
