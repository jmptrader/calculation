package calc

import "math"

// Triangle type
type Triangle struct {
	Point1, Point2, Point3 Point
}

// Area gives area.
// See http://www.mathopenref.com/coordtrianglearea.html.
func (t *Triangle) Area() float64 {
	ax := t.Point1.X
	ay := t.Point1.Y
	bx := t.Point2.X
	by := t.Point2.Y
	cx := t.Point3.X
	cy := t.Point3.Y
	return math.Abs(ax*(by-cy)+bx*(cy-ay)+cx*(ay-by)) / 2.0
}

// EqualSides returns the number of pairs of equal sides.
// 0=scalene
// 1=isosceles
// 3=equilateral
func (t *Triangle) EqualSides() int {
	// Get length of first side.
	a := t.Point1.X - t.Point2.X
	b := t.Point1.Y - t.Point2.Y
	len1 := math.Sqrt(a*a + b*b)

	// Get length of second side.
	a = t.Point2.X - t.Point3.X
	b = t.Point2.Y - t.Point3.Y
	len2 := math.Sqrt(a*a + b*b)

	// Get length of third side.
	a = t.Point1.X - t.Point3.X
	b = t.Point1.Y - t.Point3.Y
	len3 := math.Sqrt(a*a + b*b)

	// Count pairs of equal sides.
	count := 0
	if Near(len1, len2) {
		count++
	}
	if Near(len2, len3) {
		count++
	}
	if Near(len1, len3) {
		count++
	}
	return count
}

// Perim gives perimeter.
func (t *Triangle) Perim() float64 {
	// Get length of first side.
	a := t.Point1.X - t.Point2.X
	b := t.Point1.Y - t.Point2.Y
	len1 := math.Sqrt(a*a + b*b)

	// Get length of second side.
	a = t.Point2.X - t.Point3.X
	b = t.Point2.Y - t.Point3.Y
	len2 := math.Sqrt(a*a + b*b)

	// Get length of third side.
	a = t.Point1.X - t.Point3.X
	b = t.Point1.Y - t.Point3.Y
	len3 := math.Sqrt(a*a + b*b)

	// Return sum of lengths.
	return len1 + len2 + len3
}
