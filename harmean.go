package calc

// HarMean returns the harmonic mean.
func HarMean(values []float64) float64 {
	sum := 0.0
	for _, value := range values {
		sum += 1 / value
	}
	return float64(len(values)) / sum
}
