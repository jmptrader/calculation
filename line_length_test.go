package calc

import "testing"

func TestLineLength(t *testing.T) {
	// Make normal line segment.
	point1 := Point{3.0, 6.0}
	point2 := Point{6.0, 10.0}
	line := Line{point1, point2}

	// Test length of normal line segment.
	expLen := 5.0
	actLen := line.Length()
	if !Near(expLen, actLen) {
		t.Error("Expected", expLen, "got", actLen)
	}
}
