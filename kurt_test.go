package calc

import "testing"

func TestKurt(t *testing.T) {
	exp := -1.680547038
	act := Kurt([]float64{2.3, 3.5, 5.7, 7.9})
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
