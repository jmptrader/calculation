package calc

import "math"

// Length treats a line as a line segment and returns its length.
func (l *Line) Length() float64 {
	a := l.Point1.X - l.Point2.X
	b := l.Point1.Y - l.Point2.Y
	return math.Sqrt(a*a + b*b)
}
