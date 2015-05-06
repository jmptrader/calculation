package calc

import "testing"

func TestBinDist(t *testing.T) {
	exp := 0.34375
	act := BinDist(0.5, 6, 0, 2)
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
