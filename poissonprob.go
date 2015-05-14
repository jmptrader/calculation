package calc

import "math"

// PoissonProb calculates the probability value of the Poisson distribution for the mean at x.
// See https://en.wikipedia.org/wiki/Poisson_distribution.
func PoissonProb(mean float64, x int) float64 {
	res := math.Pow(math.E, -1.0*mean)
	for i := 1; i <= x; i++ {
		res *= mean / float64(i)
	}
	return res
}
