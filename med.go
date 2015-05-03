package calc

import "sort"

// Med returns the median, or middle, value.
func Med(vals []float64) float64 {
	sort.Float64s(vals)
	med := 0.0
	if len(vals)%2 == 1 {
		m := (len(vals) - 1) / 2
		med = vals[m]
	} else {
		m := len(vals)/2 - 1
		med = (vals[m] + vals[m+1]) / 2.0
	}
	return med
}
