package calc

import (
	"math"
	"testing"
)

func TestHoriz(t *testing.T) {
	// Make horizontal line.
	point1 := Point{3.0, 9.0}
	point2 := Point{5.0, 9.0}
	line := Line{point1, point2}

	// Test validity of horizontal line.
	if !line.IsValid() {
		t.Error("Expected true, got false")
	}

	// Test if horizontal line is horizontal.
	if !line.IsHoriz() {
		t.Error("Expected true, got false")
	}

	// Test if horizontal line is vertical.
	if line.IsVert() {
		t.Error("Expected false, got true")
	}

	// Test intercept of horizontal line.
	expIntc := 9.0
	actIntc := line.Intc()
	if expIntc != actIntc {
		t.Error("Expected", expIntc, "got", actIntc)
	}

	// Test slope of horizontal line.
	expSlope := 0.0
	actSlope := line.Slope()
	if expSlope != actSlope {
		t.Error("Expected", expSlope, "got", actSlope)
	}
}

func TestIsParl(t *testing.T) {
	// Make parallel lines.
	point1 := Point{2.0, 11.0}
	point2 := Point{3.0, 14.0}
	line1 := Line{point1, point2}
	point3 := Point{2.0, 16.0}
	point4 := Point{3.0, 19.0}
	line2 := Line{point3, point4}

	// Test that parallel lines are parallel.
	if !line1.IsParl(&line2) {
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
	if !line3.IsParl(&line4) {
		t.Error("Expected true, got false")
	}

	// Test non-parallel lines.
	if line1.IsParl(&line3) {
		t.Error("Expected false, got true")
	}
}

func TestIsPerp(t *testing.T) {
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

func TestIsValid(t *testing.T) {
	// Make invalid line.
	point1 := Point{3.0, 9.0}
	point2 := Point{3.0, 9.0}
	line := Line{point1, point2}

	// Test validity of invalid line.
	if line.IsValid() {
		t.Error("Expected false, got true")
	}
}

func TestLength(t *testing.T) {
	// Make normal line segment.
	point1 := Point{3.0, 6.0}
	point2 := Point{6.0, 10.0}
	line := Line{point1, point2}

	// Test length of normal line segment.
	expLen := 5.0
	actLen := line.Length()
	if expLen != actLen {
		t.Error("Expected", expLen, "got", actLen)
	}
}

func TestNormal(t *testing.T) {
	// Make normal line.
	point1 := Point{3.0, 9.0}
	point2 := Point{5.0, 13.0}
	line := Line{point1, point2}

	// Test validity of normal line.
	if !line.IsValid() {
		t.Error("Expected true, got false")
	}

	// Test if normal line is horizontal.
	if line.IsHoriz() {
		t.Error("Expected false, got true")
	}

	// Test if normal line is vertical.
	if line.IsVert() {
		t.Error("Expected false, got true")
	}

	// Test intercept of normal line.
	expIntc := 3.0
	actIntc := line.Intc()
	if expIntc != actIntc {
		t.Error("Expected", expIntc, "got", actIntc)
	}

	// Test slope of normal line.
	expSlope := 2.0
	actSlope := line.Slope()
	if expSlope != actSlope {
		t.Error("Expected", expSlope, "got", actSlope)
	}
}

func TestVert(t *testing.T) {
	// Make vertical line.
	point1 := Point{3.0, 9.0}
	point2 := Point{3.0, 13.0}
	line := Line{point1, point2}

	// Test validity of vertical line.
	if !line.IsValid() {
		t.Error("Expected true, got false")
	}

	// Test if vertical line is horizontal.
	if line.IsHoriz() {
		t.Error("Expected false, got true")
	}

	// Test if vertical line is vertical.
	if !line.IsVert() {
		t.Error("Expected true, got false")
	}

	// Test intercept of vertical line.
	actIntc := line.Intc()
	if !math.IsNaN(actIntc) {
		t.Error("Expected NaN got", actIntc)
	}

	// Test slope of vertical line.
	actSlope := line.Slope()
	if !math.IsNaN(actSlope) {
		t.Error("Expected NaN got", actSlope)
	}
}
