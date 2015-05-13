package calc

// Chebyshev returns the calculation of Chebyshev's inequality.
// In other words, k is a real number that signifies how many standard deviations from the mean, and the return value
// shows you what proportion of values lies within that many standard deviations for an arbitrary distribution.
// For example, k = 2 shows that at least 1-1/(2*2) = 1-1/4 = 75% of values lie within the +/-2 stddev limit for an
// arbitrary distribution, meaning that at most 100% - 75% = 25% lie outside the same +/-2 stddev limit.
// More accurate results can be given if more is known about the distribution, for example, the empirical rule for the
// standard distribution states that about 68% of values lie within +/-2 stddev of a normal distribution, which is a
// tighter limit.
// Note that Chebyshev is only valid when k > 1 and stddev > 1.
// See https://en.wikipedia.org/wiki/Chebyshev%27s_inequality.
func Chebyshev(k float64) float64 {
	// Test for vacuous result of probability greater than 100%.
	if k <= 1.0 {
		return 0.0
	}

	// Make Chebyshev calculation.
	return 1.0 - 1.0/(k*k)
}
