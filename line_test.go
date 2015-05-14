package calc

import (
	"math"
	"testing"
)

func TestLineHoriz(t *testing.T) {
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
	if !Near(expIntc, actIntc) {
		t.Error("Expected", expIntc, "got", actIntc)
	}

	// Test slope of horizontal line.
	expSlope := 0.0
	actSlope := line.Slope()
	if !Near(expSlope, actSlope) {
		t.Error("Expected", expSlope, "got", actSlope)
	}
}

func TestLineNormal(t *testing.T) {
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
	if !Near(expIntc, actIntc) {
		t.Error("Expected", expIntc, "got", actIntc)
	}

	// Test slope of normal line.
	expSlope := 2.0
	actSlope := line.Slope()
	if !Near(expSlope, actSlope) {
		t.Error("Expected", expSlope, "got", actSlope)
	}
}

func TestLineVert(t *testing.T) {
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
