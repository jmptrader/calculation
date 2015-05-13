package calc

import "math"

// StdErr returns the sample standard error.
// See https://en.wikipedia.org/wiki/Standard_error.
func StdErr(vals []float64) float64 {
	mean := Mean(vals)
	sum := 0.0
	for _, val := range vals {
		dev := val - mean
		sum += dev * dev
	}
	return math.Sqrt(sum / (float64(len(vals)) * float64(len(vals)-1)))
}
