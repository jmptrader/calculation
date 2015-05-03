package calc

import "testing"

func TestAdd(t *testing.T) {
	exp := Frac{29, 21}
	fact1 := Frac{2, 3}
	fact2 := Frac{5, 7}
	act1 := Add(fact1, fact2)
	if exp != act1 {
		t.Error("Expected", exp, "got", act1)
	}
	act2 := Add(fact2, fact1)
	if exp != act2 {
		t.Error("Expected", exp, "got", act2)
	}
}

func TestSub(t *testing.T) {
	exp1 := Frac{-1, 21}
	fact1 := Frac{2, 3}
	fact2 := Frac{5, 7}
	act1 := Sub(fact1, fact2)
	if exp1 != act1 {
		t.Error("Expected", exp1, "got", act1)
	}
	exp2 := Frac{1, 21}
	act2 := Sub(fact2, fact1)
	if exp2 != act2 {
		t.Error("Expected", exp2, "got", act2)
	}
}

func TestMul(t *testing.T) {
	exp := Frac{10, 21}
	fact1 := Frac{2, 3}
	fact2 := Frac{5, 7}
	act1 := Mul(fact1, fact2)
	if exp != act1 {
		t.Error("Expected", exp, "got", act1)
	}
	act2 := Mul(fact2, fact1)
	if exp != act2 {
		t.Error("Expected", exp, "got", act2)
	}
}

func TestDiv(t *testing.T) {
	exp1 := Frac{14, 15}
	fact1 := Frac{2, 3}
	fact2 := Frac{5, 7}
	act1 := Div(fact1, fact2)
	if exp1 != act1 {
		t.Error("Expected", exp1, "got", act1)
	}
	exp2 := Frac{15, 14}
	act2 := Div(fact2, fact1)
	if exp2 != act2 {
		t.Error("Expected", exp2, "got", act2)
	}
}
