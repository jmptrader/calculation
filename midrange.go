package calc

// Midrange returns the average of the min and max.
func Midrange(vals []float64) float64 {
	min := vals[0]
	max := vals[0]
	for _, val := range vals {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	return (min + max) / 2.0
}
