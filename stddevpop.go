package calc

import "math"

// StdDevPop returns the population standard deviation.
func StdDevPop(values []float64) float64 {
	mean := Mean(values)
	sum := 0.0
	for _, value := range values {
		dev := value - mean
		sum += dev * dev
	}
	return math.Sqrt(sum / float64(len(values)))
}
