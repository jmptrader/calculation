package calc

// HarMean returns the harmonic mean.
// https://en.wikipedia.org/wiki/Harmonic_mean.
func HarMean(vals []float64) float64 {
	sum := 0.0
	for _, val := range vals {
		sum += 1 / val
	}
	return float64(len(vals)) / sum
}
