package calc

import "math"

// Dist returns the distance between two coordinates.
func Dist(x1, y1, x2, y2 float64) float64 {
	a := x1 - x2
	b := y1 - y2
	return math.Sqrt(a*a + b*b)
}
