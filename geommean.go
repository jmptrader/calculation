package calc

import "math"

// GeomMean returns the geometric mean.
func GeomMean(values []float64) float64 {
	prod := 1.0
	for _, value := range values {
		prod *= value
	}
	return math.Pow(prod, 1/float64(len(values)))
}
