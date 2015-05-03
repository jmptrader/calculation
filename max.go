package calc

// Max returns the largest number.
func Max(values []float64) float64 {
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}
