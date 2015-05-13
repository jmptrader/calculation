package calc

import "math"

const Epsilon = 1e-7

// Near determines if two floating point vals are nearly equal.
func Near(x float64, y float64) bool {
	// Check if numbers are exactly equal.
	if x == y {
		return true
	}

	// Check if numbers were both near zero.
	sumAbs := math.Abs(x) + math.Abs(y)
	if sumAbs < Epsilon {
		return true
	}

	// Check if both numbers were relatively equal within tolerance.
	// Here we avoid dividing by zero because of the above.
	absDiff := math.Abs(x - y)
	return absDiff/sumAbs < 0.0000001
}
