package calc

import "sort"

// Large returns the nth largest number.
func Large(vals []float64, n int) float64 {
	sort.Float64s(vals)
	return vals[len(vals)-n]
}
