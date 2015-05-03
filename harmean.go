package calc

// HarMean returns the harmonic mean.
func HarMean(vals []float64) float64 {
	sum := 0.0
	for _, val := range vals {
		sum += 1 / val
	}
	return float64(len(vals)) / sum
}
