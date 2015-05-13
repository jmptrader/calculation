package calc

import "math"

// BinomSkew returns the moment coefficient of skewness of the binomial distribution.
// See https://en.wikipedia.org/wiki/Binomial_distribution.
func BinomSkew(n int, p float64) float64 {
	q := 1.0 - p
	return (q - p) / math.Sqrt(float64(n)*p*q)
}
