package calc

import "sort"

// FracRank returns the fraction of a list at which the given value lies.
// If the value does not exist in the list, rank will be interpolated.
func FracRank(values []float64, value float64) float64 {
	sort.Float64s(values)
	n := len(values)
	if value <= values[0] {
		return 0.0
	} else if value >= values[n-1] {
		return 1.0
	}
	for i := 0; i < n-1; i++ {
		if value == values[i] {
			return float64(i) / float64(n-1)
		} else if value < values[i+1] {
			keyFloor := float64(i) / float64(n-1)
			keyCeil := float64(i+1) / float64(n-1)
			keyDiff := keyCeil - keyFloor
			prop := (value - values[i]) / (values[i+1] - values[i])
			return keyFloor + keyDiff*prop
		}
	}
	return -1.0
}
