package calc

import "sort"

// Prod returns the product of numbers without overflow or underflow.
// Use for a series of floats that contain large numbers and small fractions.
func Prod(values []float64) float64 {
	sort.Float64s(values)
	n := len(values)
	prod := 1.0
	medium := values[n/2]
	i := 0
	j := n - 1
	for i <= j {
		if prod > medium {
			prod *= values[i]
			i++
		} else {
			prod *= values[j]
			j--
		}
	}
	return prod
}
