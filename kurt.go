package calc

// Kurt returns the sample kurtosis.
func Kurt(values []float64) float64 {
	mean := Mean(values)
	stdDev := StdDev(values)
	sum := 0.0
	for _, value := range values {
		dev := (value - mean) / stdDev
		dev *= dev
		sum += dev * dev
	}
	n := len(values)

	// Calculate left factor from its numerator and denominator.
	leftFacNumer := n * (n + 1)
	leftFacDenom := (n - 1) * (n - 2) * (n - 3)
	leftFac := float64(leftFacNumer) / float64(leftFacDenom)

	// Calculate right addend from its numerator and denominator.
	rightAdnNumer := 3 * (n - 1) * (n - 1)
	rightAdnDenom := (n - 2) * (n - 3)
	rightAdn := float64(rightAdnNumer) / float64(rightAdnDenom)

	// Return formula.
	return leftFac*sum - rightAdn
}
