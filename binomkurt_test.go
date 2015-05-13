package calc

import "testing"

func TestBinomKurt(t *testing.T) {
	exp := 2.98
	act := BinomKurt(100, 0.5)
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
