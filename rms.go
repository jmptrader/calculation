package calc

import "math"

// RMS returns the root mean square (RMS).
func RMS(vals []float64) float64 {
	sum := 0.0
	for _, val := range vals {
		sum += val * val
	}
	return math.Sqrt(sum / float64(len(vals)))
}
