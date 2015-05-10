package calc

import "testing"

func TestTrimMean(t *testing.T) {
	exp := 6.255555556
	act := TrimMean([]float64{5.2, 9.1, 9.8, 9.5, 0.4, 6.7, 5.7, 3.2, 6.7}, 0.1)
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
