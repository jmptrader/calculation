package calc

import "testing"

func TestFracLim(t *testing.T) {
	vals := []float64{2.0, 3.0, 5.0, 7.0, 11.0, 13.0, 17.0}

	// Test minimum.
	actMin := FracLim(vals, 0.0)
	if !Near(vals[0], actMin) {
		t.Error("Expected", vals[0], "got", actMin)
	}

	// Test maximum.
	actMax := FracLim(vals, 1.0)
	m := len(vals) - 1
	if !Near(vals[m], actMax) {
		t.Error("Expected", vals[m], "got", actMax)
	}

	// Test median.
	actMed := FracLim(vals, 0.5)
	m = (len(vals) - 1) / 2
	if !Near(vals[m], actMed) {
		t.Error("Expected", vals[m], "got", actMed)
	}

	// Test other.
	actOther := FracLim(vals, 0.25)
	m = (len(vals) - 1) / 4
	expOther := (vals[m] + vals[m+1]) / 2.0
	if !Near(expOther, actOther) {
		t.Error("Expected", expOther, "got", actOther)
	}

}
