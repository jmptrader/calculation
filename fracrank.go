package calc

import "sort"

// FracRank returns the fraction of a list at which the given value lies.
// If the value does not exist in the list, rank will be interpolated.
// See https://en.wikipedia.org/wiki/Percentile.
func FracRank(vals []float64, val float64) float64 {
	sort.Float64s(vals)
	n := len(vals)
	if val <= vals[0] {
		return 0.0
	} else if val >= vals[n-1] {
		return 1.0
	}
	for i := 0; i < n-1; i++ {
		if Near(val, vals[i]) {
			return float64(i) / float64(n-1)
		} else if val < vals[i+1] {
			keyFloor := float64(i) / float64(n-1)
			keyCeil := float64(i+1) / float64(n-1)
			keyDiff := keyCeil - keyFloor
			prop := (val - vals[i]) / (vals[i+1] - vals[i])
			return keyFloor + keyDiff*prop
		}
	}
	return -1.0
}
