package calc

// Min returns the smallest number.
func Min(vals []float64) float64 {
	min := vals[0]
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}
