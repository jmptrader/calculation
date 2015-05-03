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
	leftFacNumer := n * (n + 1)
	leftFacDenom := (n - 1) * (n - 2) * (n - 3)
	lf := float64(leftFacNumer) / float64(leftFacDenom)

	// Calculate right addend from its numerator and denominator.
	rightAdnNumer := 3 * (n - 1) * (n - 1)
	rightAdnDenom := (n - 2) * (n - 3)
	rightAdn := float64(rightAdnNumer) / float64(rightAdnDenom)

	// Return formula.
	return lf*sum - rightAdn
}
