package calc

import "testing"

func TestMax(t *testing.T) {
	exp := 7.9
	act := Max([]float64{5.7, 2.3, 7.9, 3.5})
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
