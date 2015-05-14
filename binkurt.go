package calc

// BinKurt returns the moment coefficient of skewness of the binomial distribution.
// See https://en.wikipedia.org/wiki/Binial_distribution.
func BinKurt(n int, p float64) float64 {
	q := 1.0 - p
	return 3.0 + (1.0-6*p*q)/(float64(n)*p*q)
}
