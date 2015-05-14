package calc

// PoissonDist calculates the probability distribution of the Poisson distribution for the mean between x1 and x2.
// See https://en.wikipedia.org/wiki/Poisson_distribution.
// @todo Implement this more efficiently to reduce redundant computations.
func PoissonDist(mean float64, x1 int, x2 int) float64 {
	sum := 0.0
	for x := x1; x <= x2; x++ {
		sum += PoissonProb(mean, x)
	}
	return sum
}
