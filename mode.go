package calc

// Mode returns the most common value(s), if any.
func Mode(vals []float64) []float64 {
	counts := make(map[float64]int)
	max := 0
	for _, val := range vals {
		counts[val]++
		if counts[val] > max {
			max = counts[val]
		}
	}
	modes := make([]float64, 1, len(counts))

	// A mode must occur more than once.
	if max > 1 {
		for val, count := range counts {
			if count == max {
				// Every val with a max count is a mode.
				modes = append(modes, val)
			}
		}
	}
	return modes
}
