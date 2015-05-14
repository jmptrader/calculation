package calc

import "testing"

func TestPoissonStdDev(t *testing.T) {
	exp := 2.236068
	act := PoissonStdDev(5)
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
