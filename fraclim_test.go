package calc

import "testing"

func TestFracLim(t *testing.T) {
	values := []float64{2.0, 3.0, 5.0, 7.0, 11.0, 13.0, 17.0}

	// Test minimum.
	actMin := FracLim(values, 0.0)
	if !Near(values[0], actMin) {
		t.Error("Expected", values[0], "got", actMin)
	}

	// Test maximum.
	actMax := FracLim(values, 1.0)
	m := len(values) - 1
	if !Near(values[m], actMax) {
		t.Error("Expected", values[m], "got", actMax)
	}

	// Test median.
	actMed := FracLim(values, 0.5)
	m = (len(values) - 1) / 2
	if !Near(values[m], actMed) {
		t.Error("Expected", values[m], "got", actMed)
	}

	// Test other.
	actOther := FracLim(values, 0.25)
	m = (len(values) - 1) / 4
	expOther := (values[m] + values[m+1]) / 2.0
	if !Near(expOther, actOther) {
		t.Error("Expected", expOther, "got", actOther)
	}

}
