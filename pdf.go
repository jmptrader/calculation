package calc

import "math"

// PDF returns phi, the probabilistic density function.
// See https://en.wikipedia.org/wiki/Normal_distribution.
func PDF(x float64) float64 {
	num := math.Pow(math.E, -0.5*x*x)
	den := math.Sqrt(2 * math.Pi)
	return num / den
}
