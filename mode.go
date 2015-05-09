package calc

import "sort"

// Mode returns the most common value(s), if any.
func Mode(vals []float64) []float64 {
	// Sort the values so we can compare for near equality.
	sort.Float64s(vals)
	cnts := make(map[float64]int)
	max := 0
	prev := 0.0
	for _, val := range vals {
		if Near(val, prev) {
			// Discard trivial differences between floating point values.
			val = prev
		} else {
			// Only save the change once a sufficient difference has occurred.
			prev = val
		}
		cnts[val]++
		if cnts[val] > max {
			max = cnts[val]
		}
	}
	modes := make([]float64, 0, len(cnts))

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
