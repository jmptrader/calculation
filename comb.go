package calc

// Comb returns the combination count of n objects taken r at a time.
// See https://en.wikipedia.org/wiki/Combination.
func Comb(n int, r int) int {
	// Since there are 2r terms, and since C(n, r) = C(n, n - r), we want to use the lesser of r and n - r.
	if n-r < r {
		r = n - r
	}

	// Get counts of all prime factors in numerator.
	nums := make(map[int]int)
	for i := n - r + 1; i <= n; i++ {
		factors := Factor(i)
		for factor, cnt := range factors {
			nums[factor] += cnt
		}
	}

	// Get counts of all prime factors in denominator.
	dens := make(map[int]int)
	for i := 2; i <= r; i++ {
		factors := Factor(i)
		for factor, cnt := range factors {
			dens[factor] += cnt
		}
	}

	// Cancel out any common factors, and calculate result.
	comb := 1
	for factor, cnt := range nums {
		f := cnt - dens[factor]
		for i := 0; i < f; i++ {
			comb *= factor
		}
	}
	return comb
}
