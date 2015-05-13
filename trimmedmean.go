package calc

import "sort"

// TrimmedMean returns the trimmed arithmetic mean.
// The highest and lowest int(len(vals) * f / 2) vals are ignored.
// See https://en.wikipedia.org/wiki/Truncated_mean.
func TrimmedMean(vals []float64, f float64) float64 {
	sort.Float64s(vals)
	low := int(float64(len(vals)) * f / 2.0)
	high := len(vals) - low
	sum := 0.0
	for i := low; i < high; i++ {
		sum += vals[i]
	}
	n := high - low
	return sum / float64(n)
}
