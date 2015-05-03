package calc

// WeightedMean returns the weighted arithmetic mean.
func WeightedMean(values []float64, weights []float64) float64 {
	sum := 0.0
	weightSum := 0.0
	for i, value := range values {
		sum += value * weights[i]
		weightSum += weights[i]
	}
	return sum / weightSum
}
