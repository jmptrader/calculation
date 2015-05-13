package calc

import "testing"

func TestBinomStdDev(t *testing.T) {
	exp := 5.0
	act := BinomStdDev(100, 0.5)
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
