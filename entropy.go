package calc

import "math"

// Entropy calculates the average amount of information in bits of a random variable.
func Entropy(p []float64) float64 {
	sum := 0.0
	for _, v := range p {
		sum += v * math.Log2(v)
	}
	return -1.0 * sum
}
