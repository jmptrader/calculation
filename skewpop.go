package calc

// SkewPop returns the population skew.
func SkewPop(values []float64) float64 {
	mean := Mean(values)
	stdDev := StdDevPop(values)
	sum := 0.0
	for _, value := range values {
		dev := (value - mean) / stdDev
		sum += dev * dev * dev
	}
	n := len(values)

	// Calculate left factor from its numerator and denominator.
	leftFacDenom := (n - 1) * (n - 2)
	leftFac := float64(n) / float64(leftFacDenom)

	// Return formula.
	return leftFac * sum
}
