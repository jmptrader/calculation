package calc

import "testing"

func TestPointDist(t *testing.T) {
	// Make points.
	point1 := Point{3.0, 6.0}
	point2 := Point{6.0, 10.0}

	// Test symmetric distance between points.
	expDist1 := 5.0
	actDist1 := point1.Dist(&point2)
	if !Near(expDist1, actDist1) {
		t.Error("Expected", expDist1, "got", actDist1)
	}
	actDist2 := point2.Dist(&point1)
	if !Near(expDist1, actDist2) {
		t.Error("Expected", expDist1, "got", actDist2)
	}

	// Test the zero distance.
	expDist2 := 0.0
	actDist3 := point1.Dist(&point1)
	if !Near(expDist2, actDist3) {
		t.Error("Expected", expDist2, "got", actDist3)
	}
}
