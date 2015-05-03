package calc

import "sort"

// Rank returns the rank of first occurrence of a number in a list in the given order.
func Rank(val float64, vals []float64, asc bool) int {
	sort.Float64s(vals)
	if asc {
		for i := 0; i < len(vals); i++ {
			if Near(val, vals[i]) {
				return i
			}
		}
	} else {
		for i := len(vals) - 1; i >= 0; i-- {
			if Near(val, vals[i]) {
				return i
			}
		}
	}
	return -1
}
