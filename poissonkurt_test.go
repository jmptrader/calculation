package calc

import "testing"

func TestPoissonKurt(t *testing.T) {
	exp := 3.2
	act := PoissonKurt(5)
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
