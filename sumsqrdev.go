package calc

// SumSqrDev returns the sum of squared deviations from the mean.
func SumSqrDev(vals []float64) float64 {
	mean := Mean(vals)
	sum := 0.0
	for _, val := range vals {
		dev := val - mean
		sum += dev * dev
	}
	return sum
}
