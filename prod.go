package calc

import "sort"

// Prod returns the product of numbers without overflow or underflow.
// Use for a series of floats that contain large numbers and small fractions.
// See https://en.wikipedia.org/wiki/Product_%28mathematics%29.
func Prod(vals []float64) float64 {
	sort.Float64s(vals)
	n := len(vals)
	prod := 1.0
	med := vals[n/2]
	i := 0
	j := n - 1
	for i <= j {
		if prod > med {
			prod *= vals[i]
			i++
		} else {
			prod *= vals[j]
			j--
		}
	}
	return prod
}
