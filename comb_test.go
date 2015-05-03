package calc

import "testing"

func TestComb(t *testing.T) {
	exp := 10
	act := Comb(5, 3)
	if exp != act {
		t.Error("Expected", exp, "got", act)
	}
}
