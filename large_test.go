package calc

import "testing"

func TestLarge(t *testing.T) {
	values := []float64{1.5, 2.5, 3.5, 4.5, 5.5}
	exp := 4.5
	act := Large(values, 2)
	if exp != act {
		t.Error("Expected", exp, "got", act)
	}
}
