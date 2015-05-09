package calc

import "testing"

func TestMin(t *testing.T) {
	exp := 2.3
	act := Min([]float64{5.7, 2.3, 7.9, 3.5})
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
