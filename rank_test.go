package calc

import "testing"

func TestRank(t *testing.T) {
	exp1 := 2
	act1 := Rank(5.7, []float64{7.9, 2.3, 5.7, 9.1}, true)
	if exp1 != act1 {
		t.Error("Expected", exp1, "got", act1)
	}
	exp2 := 3
	act2 := Rank(5.7, []float64{7.9, 2.3, 5.7, 9.1}, false)
	if exp2 != act2 {
		t.Error("Expected", exp2, "got", act2)
	}
}
