package calc

import "math"

// Intc calculates the intercept.
func (l *Line) Intc() float64 {
	if l.IsVert() {
		return math.NaN()
	}
	// Intercept is b = y - mx.
	return l.Point1.Y - l.Slope()*l.Point1.X
}
