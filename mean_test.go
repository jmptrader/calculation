package calc

import "testing"

func TestMean(t *testing.T) {
	exp := 1.5
	act := Mean([]float64{1, 2})
	if exp != act {
		t.Error("Expected", exp, "got", act)
	}
}
