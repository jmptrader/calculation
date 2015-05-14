package calc

import "math"

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
	cnt := 0
	if Near(len1, len2) {
		cnt++
	}
	if Near(len2, len3) {
		cnt++
	}
	if Near(len1, len3) {
		cnt++
	}
	return cnt
}
