package calc

import "sort"

// Prod returns the product of numbers without overflow or underflow.
// Use for a series of floats that contain large numbers and small fractions.
func Prod(vals []float64) float64 {
	sort.Float64s(vals)
	n := len(vals)
	prod := 1.0
	medium := vals[n/2]
	i := 0
	j := n - 1
	for i <= j {
		if prod > medium {
			prod *= vals[i]
			i++
		} else {
			prod *= vals[j]
			j--
		}
	}
	return prod
}
