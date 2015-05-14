package calc

import "testing"

func TestDotProd(t *testing.T) {
	exp := 3.0
	vecA := Vec([]float64{1.0, 3.0, -5.0})
	vecB := Vec([]float64{4.0, -2.0, -1.0})
	act := vecA.DotProd(vecB)
	if exp != act {
		t.Error("Expected", exp, "got", act)
	}
}
