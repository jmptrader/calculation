package calc

// Max returns the largest number.
// See https://en.wikipedia.org/wiki/Maxima_and_minima.
func Max(vals []float64) float64 {
	max := vals[0]
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}
