package calc

import "testing"

func TestDotProd(t *testing.T) {
	exp := 3.0
	vec1 := []float64{1.0, 3.0, -5.0}
	vec2 := []float64{4.0, -2.0, -1.0}
	act := DotProd(vec1, vec2)
	if exp != act {
		t.Error("Expected", exp, "got", act)
	}
}
