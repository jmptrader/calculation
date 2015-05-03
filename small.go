package calc

import "sort"

// Small returns the nth smallest number.
func Small(values []float64, n int) float64 {
	sort.Float64s(values)
	return values[n-1]
}
