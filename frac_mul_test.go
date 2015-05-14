package calc

import "testing"

func TestFracMul(t *testing.T) {
	exp := Frac{10, 21}
	frac1 := Frac{2, 3}
	frac2 := Frac{5, 7}
	act1 := Mul(frac1, frac2)
	if exp != act1 {
		t.Error("Expected", exp, "got", act1)
	}
	act2 := Mul(frac2, frac1)
	if exp != act2 {
		t.Error("Expected", exp, "got", act2)
	}
}
