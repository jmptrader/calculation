package calc

import "testing"

func TestLineIsPerp(t *testing.T) {
	// Make perpendicular lines.
	point1 := Point{3.0, 14.0}
	point2 := Point{6.0, 23.0}
	line1 := Line{point1, point2}
	point3 := Point{9.0, 4.0}
	point4 := Point{12.0, 3.0}
	line2 := Line{point3, point4}

	// Test that perpendicular lines are perpendicular.
	if !line1.IsPerp(&line2) {
		t.Error("Expected true, got false")
	}

	// Make horizontal and vertial lines.
	point5 := Point{3.0, 9.0}
	point6 := Point{3.0, 13.0}
	line3 := Line{point5, point6}
	point7 := Point{10.0, 5.0}
	point8 := Point{20.0, 5.0}
	line4 := Line{point7, point8}

	// Test that vertical lines are perpendicular.
	if !line3.IsPerp(&line4) {
		t.Error("Expected true, got false")
	}

	// Test non-perpendicular lines.
	if line1.IsPerp(&line3) {
		t.Error("Expected false, got true")
	}
}
