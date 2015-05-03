package calc

import "sort"

// Med returns the median, or middle, value.
func Med(values []float64) float64 {
	sort.Float64s(values)
	med := 0.0
	if len(values)%2 == 1 {
		m := (len(values) - 1) / 2
		med = values[m]
	} else {
		m := len(values)/2 - 1
		med = (values[m] + values[m+1]) / 2.0
	}
	return med
}
