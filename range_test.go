package calc

import "testing"

func TestRange(t *testing.T) {
	exp := 5.6
	act := Range([]float64{2.3, 3.5, 5.7, 7.9})
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
