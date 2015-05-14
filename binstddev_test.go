package calc

import "testing"

func TestBinStdDev(t *testing.T) {
	exp := 5.0
	act := BinStdDev(100, 0.5)
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
