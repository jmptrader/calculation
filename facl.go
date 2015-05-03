package calc

import "strconv"

// Facl returns a factorial lookup table for n!.
// Results are limited to 12! for 32-bit and 20! for 64-bit machines.
func Facl() []int {
	var lim int
	if strconv.IntSize == 64 {
		lim = 21
	} else {
		lim = 13
	}
	facls := make([]int, lim)
	facls[0] = 1
	for i := 1; i < lim; i++ {
		facls[i] = facls[i-1] * i
	}
	return facls
}
