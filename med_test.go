package calc

import "testing"

func TestMed(t *testing.T) {
	exp1 := 5.7
	act1 := Med([]float64{7.9, 2.3, 5.7})
	if !Near(exp1, act1) {
		t.Error("Expected", exp1, "got", act1)
	}
	exp2 := 6.8
	act2 := Med([]float64{7.9, 2.3, 5.7, 9.1})
	if !Near(exp2, act2) {
		t.Error("Expected", exp2, "got", act2)
	}
}
