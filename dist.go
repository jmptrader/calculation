package calc

import "math"

// Dist returns the distance between two coordinates.
// See https://en.wikipedia.org/wiki/Distance.
func Dist(point1 []float64, point2 []float64) float64 {
	sum := 0.0
	for i, _ := range point1 {
		term := point2[i] - point1[i]
		sum += term * term
	}
	return math.Sqrt(sum)
}
