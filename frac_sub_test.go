package calc

import "testing"

func TestFracSub(t *testing.T) {
	exp1 := Frac{-1, 21}
	frac1 := Frac{2, 3}
	frac2 := Frac{5, 7}
	act1 := Sub(frac1, frac2)
	if exp1 != act1 {
		t.Error("Expected", exp1, "got", act1)
	}
	exp2 := Frac{1, 21}
	act2 := Sub(frac2, frac1)
	if exp2 != act2 {
		t.Error("Expected", exp2, "got", act2)
	}
}
