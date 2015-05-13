package calc

import "math"

// Gaussian returns the normal distribution.
// See https://en.wikipedia.org/wiki/Normal_distribution.
func Gaussian(x float64, mean float64, stdDev float64) float64 {
	// Calculate exponents.
	exp1 := x - mean
	exp1 *= exp1
	exp2 := stdDev * stdDev

	// Calculate numerator and denominator.
	num := math.Pow(math.E, -0.5*exp1/exp2)
	den := stdDev * math.Sqrt(2*math.Pi)
	return num / den
}
