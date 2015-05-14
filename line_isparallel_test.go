package calc

import "testing"

func TestLineIsParallel(t *testing.T) {
	// Make parallel lines.
	point1 := Point{2.0, 11.0}
	point2 := Point{3.0, 14.0}
	line1 := Line{point1, point2}
	point3 := Point{2.0, 16.0}
	point4 := Point{3.0, 19.0}
	line2 := Line{point3, point4}

	// Test that parallel lines are parallel.
	if !line1.IsParallel(&line2) {
		t.Error("Expected true, got false")
	}

	// Make vertical lines.
	point5 := Point{3.0, 9.0}
	point6 := Point{3.0, 13.0}
	line3 := Line{point5, point6}
	point7 := Point{3.0, 9.0}
	point8 := Point{3.0, 13.0}
	line4 := Line{point7, point8}

	// Test that vertical lines are parallel.
	if !line3.IsParallel(&line4) {
		t.Error("Expected true, got false")
	}

	// Test non-parallel lines.
	if line1.IsParallel(&line3) {
		t.Error("Expected false, got true")
	}
}
