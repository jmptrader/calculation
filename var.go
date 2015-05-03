package calc

// Var returns the sample variance.
func Var(vals []float64) float64 {
	mean := Mean(vals)
	sum := 0.0
	for _, val := range vals {
		dev := val - mean
		sum += dev * dev
	}
	return sum / float64(len(vals)-1)
}
