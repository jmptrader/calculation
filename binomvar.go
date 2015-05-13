package calc

// BinomVar returns the variance of the binomial distribution.
// See https://en.wikipedia.org/wiki/Binomial_distribution.
func BinomVar(n int, p float64) float64 {
	q := 1.0 - p
	return float64(n) * p * q
}
