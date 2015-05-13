package calc

import "math"

// RMS returns the root mean square (RMS).
// See https://en.wikipedia.org/wiki/Root_mean_square.
func RMS(vals []float64) float64 {
	sum := 0.0
	for _, val := range vals {
		sum += val * val
	}
	return math.Sqrt(sum / float64(len(vals)))
}
