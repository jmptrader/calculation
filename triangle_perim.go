package calc

import "math"

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
