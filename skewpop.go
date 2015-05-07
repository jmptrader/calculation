package calc

// SkewPop returns the population skew.
func SkewPop(vals []float64) float64 {
	mean := Mean(vals)
	stdDev := StdDevPop(vals)
	sum := 0.0
	for _, val := range vals {
		dev := (val - mean) / stdDev
		sum += dev * dev * dev
	}
	n := len(vals)

	// Calculate left factor from its numerator and denominator.
	leftFactorDen := (n - 1) * (n - 2)
	leftFactor := float64(n) / float64(leftFactorDen)

	// Return formula.
	return leftFactor * sum
}
