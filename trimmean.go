package calc

import "sort"

// TrimMean returns the trimmed arithmetic mean.
// The highest and lowest int(len(vals) * f / 2) vals are ignored.
func TrimMean(vals []float64, f float64) float64 {
	sort.Float64s(vals)
	low := int(float64(len(vals)) * f / 2.0)
	high := len(vals) - low
	sum := 0.0
	for i := low; i <= high; i++ {
		sum += vals[i]
	}
	n := high - low
	return sum / float64(n)
}
