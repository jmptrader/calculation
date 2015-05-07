package calc

// KurtPop returns the population kurtosis.
func KurtPop(vals []float64) float64 {
	mean := Mean(vals)
	stdDev := StdDevPop(vals)
	sum := 0.0
	for _, val := range vals {
		dev := (val - mean) / stdDev
		dev *= dev
		sum += dev * dev
	}
	n := len(vals)

	// Calculate left factor from its numerator and denominator.
	leftFactorNum := n * (n + 1)
	leftFactorDen := (n - 1) * (n - 2) * (n - 3)
	lf := float64(leftFactorNum) / float64(leftFactorDen)

	// Calculate right addend from its numerator and denominator.
	rightAddendNum := 3 * (n - 1) * (n - 1)
	rightAddendDen := (n - 2) * (n - 3)
	rightAddend := float64(rightAddendNum) / float64(rightAddendDen)

	// Return formula.
	return lf*sum - rightAddend
}
