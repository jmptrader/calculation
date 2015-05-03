package calc

// SumSqrDev returns the sum of squared deviations from the mean.
func SumSqrDev(values []float64) float64 {
	mean := Mean(values)
	sum := 0.0
	for _, value := range values {
		dev := value - mean
		sum += dev * dev
	}
	return sum
}
