package calc

import "testing"

func TestFracRank(t *testing.T) {
	values := []float64{1.5, 2.5, 3.5, 4.5, 5.5}
	exp1 := 0.25
	act1 := FracRank(values, 2.5)
	if exp1 != act1 {
		t.Error("Expected", exp1, "got", act1)
	}
	exp2 := 0.375
	act2 := FracRank(values, 3.0)
	if exp2 != act2 {
		t.Error("Expected", exp2, "got", act2)
	}
}
