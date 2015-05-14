package calc

import "testing"

func TestBinVar(t *testing.T) {
	exp := 25.0
	act := BinVar(100, 0.5)
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
