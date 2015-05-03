package calc

import "math"

// Moment returns the rth moment about a point.
func Moment(vals []float64, point float64, r int) float64 {
	sum := 0.0
	for _, val := range vals {
		sum += math.Pow(val-point, float64(r))
	}
	return sum / float64(len(vals))
}
