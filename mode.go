package calc

// Mode returns the most common value(s), if any.
func Mode(values []float64) []float64 {
	counts := make(map[float64]int)
	max := 0
	for _, value := range values {
		counts[value]++
		if counts[value] > max {
			max = counts[value]
		}
	}
	modes := make([]float64, 1, len(counts))

	// A mode must occur more than once.
	if max > 1 {
		for value, count := range counts {
			if count == max {
				// Every value with a max count is a mode.
				modes = append(modes, value)
			}
		}
	}
	return modes
}
