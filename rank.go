package calc

import "sort"

// Rank returns the rank of first occurrence of a number in a list in the given order.
func Rank(value float64, values []float64, ascending bool) int {
	sort.Float64s(values)
	if ascending {
		for i := 0; i < len(values); i++ {
			if Near(value, values[i]) {
				return i
			}
		}
	} else {
		for i := len(values) - 1; i >= 0; i-- {
			if Near(value, values[i]) {
				return i
			}
		}
	}
	return -1
}
