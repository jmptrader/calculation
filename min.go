package calc

// Min returns the smallest number.
func Min(values []float64) float64 {
	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min
}
