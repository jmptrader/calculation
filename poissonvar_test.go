package calc

import "testing"

func TestPoissonVar(t *testing.T) {
	exp := 5.0
	act := PoissonVar(5)
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
