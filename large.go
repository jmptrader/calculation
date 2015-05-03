package calc

import "sort"

// Large returns the nth largest number.
func Large(values []float64, n int) float64 {
	sort.Float64s(values)
	return values[len(values)-n]
}
