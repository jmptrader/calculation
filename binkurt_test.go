package calc

import "testing"

func TestBinKurt(t *testing.T) {
	exp := 2.98
	act := BinKurt(100, 0.5)
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
