package calc

import (
	"math"
	"testing"
)

func TestTriangleEquilateral(t *testing.T) {
	// Test an equilateral triangle.
	point1 := Point{2.0, 2.0}
	point2 := Point{6.0, 2.0}
	point3 := Point{4.0, 2.0 + 2.0*math.Sqrt(3.0)}
	triangle := Triangle{point1, point2, point3}
	expArea := 4 * math.Sqrt(3)
	actArea := triangle.Area()
	if !Near(expArea, actArea) {
		t.Error("Expected", expArea, "got", actArea)
	}
	expPerim := 12.0
	actPerim := triangle.Perim()
	if !Near(expPerim, actPerim) {
		t.Error("Expected", expPerim, "got", actPerim)
	}
	expEqualSides := 3
	actEqualSides := triangle.EqualSides()
	if expEqualSides != actEqualSides {
		t.Error("Expected", expEqualSides, "got", actEqualSides)
	}
}

func TestTriangleRight(t *testing.T) {
	// Test a 3-4-5 right triangle.
	point1 := Point{2.0, 2.0}
	point2 := Point{2.0, 6.0}
	point3 := Point{5.0, 2.0}
	triangle := Triangle{point1, point2, point3}
	expArea := 6.0
	actArea := triangle.Area()
	if !Near(expArea, actArea) {
		t.Error("Expected", expArea, "got", actArea)
	}
	expPerim := 12.0
	actPerim := triangle.Perim()
	if !Near(expPerim, actPerim) {
		t.Error("Expected", expPerim, "got", actPerim)
	}
	expEqualSides := 0
	actEqualSides := triangle.EqualSides()
	if expEqualSides != actEqualSides {
		t.Error("Expected", expEqualSides, "got", actEqualSides)
	}
}
