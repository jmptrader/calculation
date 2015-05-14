package calc

import "testing"

func TestLineIsValid(t *testing.T) {
	// Make invalid line.
	point1 := Point{3.0, 9.0}
	point2 := Point{3.0, 9.0}
	line := Line{point1, point2}

	// Test validity of invalid line.
	if line.IsValid() {
		t.Error("Expected false, got true")
	}
}
