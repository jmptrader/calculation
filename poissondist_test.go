package calc

import "testing"

func TestPoissonDist(t *testing.T) {
	exp := 0.26502592
	act := PoissonDist(5, 0, 3)
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
