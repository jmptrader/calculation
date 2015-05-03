package calc

import "sort"

// Small returns the nth smallest number.
func Small(vals []float64, n int) float64 {
	sort.Float64s(vals)
	return vals[n-1]
}
