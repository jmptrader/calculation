package calc

import "testing"

func TestAdd(t *testing.T) {
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

func TestSub(t *testing.T) {
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

func TestMul(t *testing.T) {
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

func TestDiv(t *testing.T) {
	exp1 := Frac{14, 15}
	frac1 := Frac{2, 3}
	frac2 := Frac{5, 7}
	act1 := Div(frac1, frac2)
	if exp1 != act1 {
		t.Error("Expected", exp1, "got", act1)
	}
	exp2 := Frac{15, 14}
	act2 := Div(frac2, frac1)
	if exp2 != act2 {
		t.Error("Expected", exp2, "got", act2)
	}
}
