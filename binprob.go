package calc

// BinProb returns the probablility of k successes in n trials in a binomial distribution.
// See https://en.wikipedia.org/wiki/Binomial_distribution.
func BinProb(prob float64, n int, k int) float64 {
	factors := make([]float64, 0, 3*n)
	for i := 1; i <= k; i++ {
		// prob^k * 1/i! * n(n-1)...(n-k+1)
		factors = append(factors, prob, 1.0/float64(i), float64(n-i+1))
	}
	negProb := 1.0 - prob
	max := n - k
	for i := 1; i <= max; i++ {
		// (1-prob)^(n-k)
		factors = append(factors, negProb)
	}
	return Prod(factors)
}
