package calc

import "testing"

func TestSkew(t *testing.T) {
	exp := 0.43826956
	act := Skew([]float64{2.3, 3.5, 5.7, 7.9})
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
