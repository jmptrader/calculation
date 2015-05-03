package calc

import "math"

// Near determines if two floating point values are nearly equal.
func Near(x float64, y float64) bool {
	if x == y {
		return true
	}
	return math.Abs((x-y)/x) < 0.0000001
}
