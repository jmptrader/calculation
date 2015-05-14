package calc

// BinMean returns the mean of the binomial distribution.
// See https://en.wikipedia.org/wiki/Binial_distribution.
func BinMean(n int, p float64) float64 {
	return float64(n) * p
}
