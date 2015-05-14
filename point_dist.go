package calc

import "math"

// Dist returns the distance between two points.
func (p1 *Point) Dist(p2 *Point) float64 {
	a := p1.X - p2.X
	b := p1.Y - p2.Y
	return math.Sqrt(a*a + b*b)
}