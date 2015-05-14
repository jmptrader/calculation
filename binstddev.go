package calc

import "math"

// BinStdDev returns the standard deviation of the binomial distribution.
// See https://en.wikipedia.org/wiki/Binial_distribution.
func BinStdDev(n int, p float64) float64 {
	q := 1.0 - p
	return math.Sqrt(float64(n) * p * q)
}
