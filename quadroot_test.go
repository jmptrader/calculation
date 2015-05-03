package calc

import "testing"

func TestQuadRoot0(t *testing.T) {
	act := QuadRoot(1.0, 1.0, 1.0)
	if len(act) != 0 {
		t.Error("Expected 0 roots, got", len(act))
	}
}
func TestQuadRoot1(t *testing.T) {
	exp := []float64{3.0}
	act := QuadRoot(1.0, -6.0, 9.0)
	if len(act) != 1 || act[0] != exp[0] {
		t.Error("Expected", exp, "got", act)
	}
}
func TestQuadRoot2(t *testing.T) {
	exp := []float64{3.0, -2.0}
	act := QuadRoot(1.0, -1.0, -6.0)
	if len(act) != 2 || act[0] != exp[0] || act[1] != exp[1] {
		t.Error("Expected", exp, "got", act)
	}
}
