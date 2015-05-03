package calc

// Var returns the sample variance.
func Var(values []float64) float64 {
	mean := Mean(values)
	sum := 0.0
	for _, value := range values {
		dev := value - mean
		sum += dev * dev
	}
	return sum / float64(len(values)-1)
}
