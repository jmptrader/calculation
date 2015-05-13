package calc

// Factor returns the counts of prime factors of an int.
// See https://en.wikipedia.org/wiki/Prime_factor.
func Factor(n int) map[int]int {
	factors := make(map[int]int)
	factor := 2
	for n != 1 {
		if n%factor == 0 {
			n /= factor
			factors[factor]++
		} else {
			factor++
		}
	}
	return factors
}
