package calc

import "testing"

func TestPoissonSkew(t *testing.T) {
	exp := 0.4472136
	act := PoissonSkew(5)
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
