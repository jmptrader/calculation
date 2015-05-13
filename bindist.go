package calc

// BinDist returns the probability for a binomial distribution between k1 and k2 trials.
// See https://en.wikipedia.org/wiki/Binomial_distribution.
func BinDist(prob float64, n int, k1 int, k2 int) float64 {
	sum := 0.0
	for k := k1; k <= k2; k++ {
		sum += BinProb(prob, n, k)
	}
	return sum
}
