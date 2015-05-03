package calc

// VarPop returns the population variance.
func VarPop(values []float64) float64 {
	mean := Mean(values)
	sum := 0.0
	for _, value := range values {
		dev := value - mean
		sum += dev * dev
	}
	return sum / float64(len(values))
}
