package calc

// Range returns the difference between max and min.
func Range(values []float64) float64 {
	min := values[0]
	max := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return max - min
}
