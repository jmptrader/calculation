package calc

import "math"

// PoissonStdDev returns the standard deviation of the Poisson distribution.
// See https://en.wikipedia.org/wiki/Poisson_distribution.
func PoissonStdDev(mean int) float64 {
	return math.Sqrt(float64(mean))
}
