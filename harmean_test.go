package calc

import "testing"

func TestHarMean(t *testing.T) {
	exp := 3.91191246
	act := HarMean([]float64{2.3, 3.5, 5.7, 7.9})
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
