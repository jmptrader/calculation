package calc

import "testing"

func TestPoissonMean(t *testing.T) {
	exp := 5.0
	act := PoissonMean(5)
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
