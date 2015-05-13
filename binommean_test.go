package calc

import "testing"

func TestBinomMean(t *testing.T) {
	exp := 50.0
	act := BinomMean(100, 0.5)
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
