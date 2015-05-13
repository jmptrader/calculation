package calc

import "testing"

func TestBinomVar(t *testing.T) {
	exp := 25.0
	act := BinomVar(100, 0.5)
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
