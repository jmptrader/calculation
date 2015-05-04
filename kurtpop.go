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
	leftFacNum := n * (n + 1)
	leftFacDen := (n - 1) * (n - 2) * (n - 3)
	lf := float64(leftFacNum) / float64(leftFacDen)

	// Calculate right addend from its numerator and denominator.
	rightAdnNum := 3 * (n - 1) * (n - 1)
	rightAdnDen := (n - 2) * (n - 3)
	rightAdn := float64(rightAdnNum) / float64(rightAdnDen)

	// Return formula.
	return lf*sum - rightAdn
}
