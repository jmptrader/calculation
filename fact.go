package calc

import "strconv"

// Fact returns a factorial lookup table for n!.
// Results are limited to 12! for 32-bit and 20! for 64-bit machines.
// See https://en.wikipedia.org/wiki/Factorial.
func Fact() []int {
	var lim int
	if strconv.IntSize == 64 {
		lim = 21
	} else {
		lim = 13
	}
	facts := make([]int, lim)
	facts[0] = 1
	for i := 1; i < lim; i++ {
		facts[i] = facts[i-1] * i
	}
	return facts
}
