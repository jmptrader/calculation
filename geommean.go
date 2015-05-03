package calc

import "math"

// GeomMean returns the geometric mean.
func GeomMean(vals []float64) float64 {
	prod := 1.0
	for _, val := range vals {
		prod *= val
	}
	return math.Pow(prod, 1/float64(len(vals)))
}
