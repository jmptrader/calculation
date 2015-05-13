package calc

// BinomMean returns the mean of the binomial distribution.
// See https://en.wikipedia.org/wiki/Binomial_distribution.
func BinomMean(n int, p float64) float64 {
	return float64(n) * p
}
