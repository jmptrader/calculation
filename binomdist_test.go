package calc

import "testing"

func TestBinomDist(t *testing.T) {
	exp := 0.34375
	act := BinomDist(0.5, 6, 0, 2)
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
