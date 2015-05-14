package calc

import "testing"

func TestBinSkew(t *testing.T) {
	exp := 0.0
	act := BinSkew(100, 0.5)
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
