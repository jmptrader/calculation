package calc

import "testing"

func TestBinomProb(t *testing.T) {
	exp := 0.234375
	act := BinomProb(0.5, 6, 2)
	if exp != act {
		t.Error("Expected", exp, "got", act)
	}
}
