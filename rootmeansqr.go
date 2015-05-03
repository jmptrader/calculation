package calc

import "math"

// RootMeanSqr returns the root mean square (RMS).
func RootMeanSqr(values []float64) float64 {
	sum := 0.0
	for _, value := range values {
		sum += value * value
	}
	return math.Sqrt(sum / float64(len(values)))
}
