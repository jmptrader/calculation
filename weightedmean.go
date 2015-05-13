package calc

// WeightedMean returns the weighted arithmetic mean.
// See https://en.wikipedia.org/wiki/Weighted_arithmetic_mean.
func WeightedMean(vals []float64, weights []float64) float64 {
	sum := 0.0
	weightSum := 0.0
	for i, val := range vals {
		sum += val * weights[i]
		weightSum += weights[i]
	}
	return sum / weightSum
}
