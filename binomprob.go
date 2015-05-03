package calc

// BinomProb returns the probablility of k successes in n trials in a binomial distribution.
func BinomProb(prob float64, n int, k int) float64 {
	facs := make([]float64, 0, 3*n)
	for i := 1; i <= k; i++ {
		// prob^k * 1/i! * n(n-1)...(n-k+1)
		facs = append(facs, prob, 1.0/float64(i), float64(n-i+1))
	}
	negProb := 1.0 - prob
	max := n - k
	for i := 1; i <= max; i++ {
		// (1-prob)^(n-k)
		facs = append(facs, negProb)
	}
	return Prod(facs)
}
