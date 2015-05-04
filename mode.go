package calc

// Mode returns the most common value(s), if any.
func Mode(vals []float64) []float64 {
	cnts := make(map[float64]int)
	max := 0
	for _, val := range vals {
		cnts[val]++
		if cnts[val] > max {
			max = cnts[val]
		}
	}
	modes := make([]float64, 1, len(cnts))

	// A mode must occur more than once.
	if max > 1 {
		for val, cnt := range cnts {
			if cnt == max {
				// Every val with a max count is a mode.
				modes = append(modes, val)
			}
		}
	}
	return modes
}
