package calc

// VarPop returns the population variance.
// See https://en.wikipedia.org/wiki/Variance.
func VarPop(vals []float64) float64 {
	mean := Mean(vals)
	sum := 0.0
	for _, val := range vals {
		dev := val - mean
		sum += dev * dev
	}
	return sum / float64(len(vals))
}
