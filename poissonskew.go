package calc

import "math"

// PoissonSkew returns the moment coefficient of skewness of the Poisson distribution.
// See https://en.wikipedia.org/wiki/Poisson_distribution.
func PoissonSkew(mean int) float64 {
	return 1.0 / math.Sqrt(float64(mean))
}
