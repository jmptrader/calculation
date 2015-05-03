package calc

import "testing"

func TestProd(t *testing.T) {
	exp := 120.0
	act := Prod([]float64{1.0, 2.0, 3.0, 4.0, 5.0})
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
