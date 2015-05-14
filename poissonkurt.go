package calc

// PoissonKurt returns the moment coefficient of kurtosis of the Poisson distribution.
// See https://en.wikipedia.org/wiki/Poisson_distribution.
func PoissonKurt(mean int) float64 {
	return 3.0 + 1.0/float64(mean)
}
