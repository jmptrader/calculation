package calc

import "math"

// MeanAbsDev returns the arithmetic mean of the absolute deviations.
func MeanAbsDev(vals []float64) float64 {
	mean := Mean(vals)
	sum := 0.0
	for _, val := range vals {
		sum += math.Abs(val - mean)
	}
	return sum / float64(len(vals))
}
