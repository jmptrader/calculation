package calc

import "math"

// Slope calculates the slope.
func (l *Line) Slope() float64 {
	if l.IsVert() {
		return math.NaN()
	}
	// Slope is rise over run = (y2 - y1) / (x2 - x1).
	return (l.Point2.Y - l.Point1.Y) / (l.Point2.X - l.Point1.X)
}
