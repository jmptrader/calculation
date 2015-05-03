package calc

import "math"

// Line type
type Line struct {
	Point1, Point2 Point
}

// Intc calculates the intercept.
func (l *Line) Intc() float64 {
	if l.IsVert() {
		return math.NaN()
	}
	// Intercept is b = y - mx.
	return l.Point1.Y - l.Slope()*l.Point1.X
}

// IsHoriz checks if a line is horizontal.
func (l *Line) IsHoriz() bool {
	return Near(l.Point1.Y, l.Point2.Y)
}

// IsParl checks if l1 is parallel to l2.
func (l1 *Line) IsParl(l2 *Line) bool {
	if l1.IsVert() && l2.IsVert() {
		return true
	}
	return Near(l1.Slope(), l2.Slope())
}

// IsPerp checks if l1 is perpendicular to l2.
func (l1 *Line) IsPerp(l2 *Line) bool {
	if l1.IsVert() && l2.IsHoriz() {
		return true
	}
	if l1.IsHoriz() && l2.IsVert() {
		return true
	}
	return Near(l1.Slope()*l2.Slope(), -1.0)
}

// IsValid checks if the line is a valid line definition of two different points.
func (l *Line) IsValid() bool {
	return !Near(l.Point1.X, l.Point2.X) || !Near(l.Point1.Y, l.Point2.Y)
}

// IsVert checks if a line is vertical.
func (l *Line) IsVert() bool {
	return Near(l.Point1.X, l.Point2.X)
}

// Length treats a line as a line segment and returns its length.
func (l *Line) Length() float64 {
	a := l.Point1.X - l.Point2.X
	b := l.Point1.Y - l.Point2.Y
	return math.Sqrt(a*a + b*b)
}

// Slope calculates the slope.
func (l *Line) Slope() float64 {
	if l.IsVert() {
		return math.NaN()
	}
	// Slope is rise over run = (y2 - y1) / (x2 - x1).
	return (l.Point2.Y - l.Point1.Y) / (l.Point2.X - l.Point1.X)
}
