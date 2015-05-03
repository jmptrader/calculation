package calc

import "sort"

// TrimMean returns the trimmed arithmetic mean.
// The highest and lowest int(len(values) * f / 2) values are ignored.
func TrimMean(values []float64, f float64) float64 {
	sort.Float64s(values)
	low := int(float64(len(values)) * f / 2.0)
	high := len(values) - low
	sum := 0.0
	for i := low; i <= high; i++ {
		sum += values[i]
	}
	n := high - low
	return sum / float64(n)
}
