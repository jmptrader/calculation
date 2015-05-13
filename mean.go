package calc

// Mean returns the arithmetic mean.
// See https://en.wikipedia.org/wiki/Arithmetic_mean.
func Mean(vals []float64) float64 {
	sum := 0.0
	for _, val := range vals {
		sum += val
	}
	return sum / float64(len(vals))
}
