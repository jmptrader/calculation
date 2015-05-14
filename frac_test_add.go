package calc

import "testing"

func TestFracAdd(t *testing.T) {
	exp := Frac{29, 21}
	frac1 := Frac{2, 3}
	frac2 := Frac{5, 7}
	act1 := Add(frac1, frac2)
	if exp != act1 {
		t.Error("Expected", exp, "got", act1)
	}
	act2 := Add(frac2, frac1)
	if exp != act2 {
		t.Error("Expected", exp, "got", act2)
	}
}
