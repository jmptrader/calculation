package calc

// Range returns the difference between max and min.
func Range(vals []float64) float64 {
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
	return max - min
}
