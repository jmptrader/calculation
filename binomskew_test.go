package calc

import "testing"

func TestBinomSkew(t *testing.T) {
	exp := 0.0
	act := BinomSkew(100, 0.5)
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
