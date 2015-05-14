package calc

import "math"

// Poisson calculates a value of the Poisson distribution for x and the mean.
// See https://en.wikipedia.org/wiki/Poisson_distribution.
func Poisson(x int, mean float64) float64 {
	res := math.Pow(math.E, -1.0*mean)
	for i := 1; i <= x; i++ {
		res *= mean / float64(i)
	}
	return res
}
