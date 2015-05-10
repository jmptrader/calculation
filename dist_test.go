package calc

import "testing"

func TestDist(t *testing.T) {
	// Test symmetric distance between coordinates.
	expDist1 := 5.0
	point1 := []float64{3.0, 6.0}
	point2 := []float64{6.0, 10.0}
	actDist1 := Dist(point1, point2)
	if !Near(expDist1, actDist1) {
		t.Error("Expected", expDist1, "got", actDist1)
	}
	actDist2 := Dist(point2, point1)
	if !Near(expDist1, actDist2) {
		t.Error("Expected", expDist1, "got", actDist2)
	}

	// Test the zero distance.
	expDist2 := 0.0
	actDist3 := Dist(point1, point1)
	if !Near(expDist2, actDist3) {
		t.Error("Expected", expDist2, "got", actDist3)
	}
}
