package calc

import "math"

// StdDevPop returns the population standard deviation.
func StdDevPop(vals []float64) float64 {
	mean := Mean(vals)
	sum := 0.0
	for _, val := range vals {
		dev := val - mean
		sum += dev * dev
	}
	return math.Sqrt(sum / float64(len(vals)))
}
