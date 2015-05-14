package calc

import "math"

// BinSkew returns the moment coefficient of skewness of the binomial distribution.
// See https://en.wikipedia.org/wiki/Binial_distribution.
func BinSkew(n int, p float64) float64 {
	q := 1.0 - p
	return (q - p) / math.Sqrt(float64(n)*p*q)
}
