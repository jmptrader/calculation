package calc

import "math"

// Stirling calculates Stirling's approximation to n!.
func Stirling(n int) float64 {
	f := float64(n)
	term1 := math.Sqrt(2.0 * math.Pi * f)
	term2 := math.Pow(f, f)
	term3 := math.Pow(math.E, -1*f)
	return term1 * term2 * term3
}
