package calc

import "math"

// Moment returns the rth moment about a point.
func Moment(values []float64, point float64, r int) float64 {
	sum := 0.0
	for _, value := range values {
		sum += math.Pow(value-point, float64(r))
	}
	return sum / float64(len(values))
}
