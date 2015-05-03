package calc

// BinomDist returns the probability for a binomial distribution between k1 and k2 trials.
func BinomDist(prob float64, n int, k1 int, k2 int) float64 {
	sum := 0.0
	for k := k1; k <= k2; k++ {
		sum += BinomProb(prob, n, k)
	}
	return sum
}
