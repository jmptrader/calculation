package calc

// Mean returns the arithmetic mean.
func Mean(vals []float64) float64 {
	sum := 0.0
	for _, val := range vals {
		sum += val
	}
	return sum / float64(len(vals))
}
