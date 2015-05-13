package calc

import "testing"

func TestEntropy(t *testing.T) {
	probs := []float64{0.25, 0.25, 0.25, 0.25}
	act := Entropy(probs)
	if !Near(act, 2.0) {
		t.Error("Expected 2.0, got ", act)
	}
}
