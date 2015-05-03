package calc

import "math"

// MeanAbsDev returns the arithmetic mean of the absolute deviations.
func MeanAbsDev(values []float64) float64 {
	mean := Mean(values)
	sum := 0.0
	for _, value := range values {
		sum += math.Abs(value - mean)
	}
	return sum / float64(len(values))
}
