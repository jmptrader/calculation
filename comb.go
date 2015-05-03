package calc

// Comb returns the combination count of n objects taken r at a time.
func Comb(n int, r int) int {
	// Since there are 2r terms, and since C(n, r) = C(n, n - r), we want to use the lesser of r and n - r.
	if n-r < r {
		r = n - r
	}

	// Get counts of all prime factors in numerator.
	numers := make(map[int]int)
	for i := n - r + 1; i <= n; i++ {
		facs := Fac(i)
		for fac, count := range facs {
			numers[fac] += count
		}
	}

	// Get counts of all prime factors in denominator.
	denoms := make(map[int]int)
	for i := 2; i <= r; i++ {
		facs := Fac(i)
		for fac, count := range facs {
			denoms[fac] += count
		}
	}

	// Cancel out any common factors, and calculate result.
	comb := 1
	for fac, count := range numers {
		f := count - denoms[fac]
		for i := 0; i < f; i++ {
			comb *= fac
		}
	}
	return comb
}
