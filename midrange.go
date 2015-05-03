package calc

// Midrange returns the average of the min and max.
func Midrange(values []float64) float64 {
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
	return (min + max) / 2.0
}
