package calc

import "testing"

func TestMode(t *testing.T) {
	act1 := Mode([]float64{2.3, 3.5, 5.7, 7.9})
	if len(act1) > 0 {
		t.Error("Expected no mode, got mode")
	}
	act2 := Mode([]float64{2.3, 3.5, 5.7, 3.50000001, 7.9})
	exp2 := 3.5
	if len(act2) != 1 || !Near(act2[0], exp2) {
		t.Error("Expected", exp2, "got", act2[0])
	}
	act3 := Mode([]float64{2.29999999, 2.3, 3.5, 5.7, 3.50000001, 7.9})
	if len(act3) != 2 {
		t.Error("Expected two modes, got", len(act3))
	}
}
