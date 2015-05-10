package calc

import "sort"

// Rank returns the rank of first occurrence of a number in a list in the given order.
func Rank(num float64, vals []float64, asc bool) int {
	sort.Float64s(vals)
	for i, val := range vals {
		if Near(val, num) {
			if asc {
				return i + 1
			} else {
				return len(vals) - i
			}
		}
	}

	return -1
}
