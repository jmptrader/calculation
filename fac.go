package calc

// Fac returns the counts of prime factors of an int.
func Fac(n int) map[int]int {
	facs := make(map[int]int)
	fac := 2
	for n != 1 {
		if n%fac == 0 {
			n /= fac
			facs[fac]++
		} else {
			fac++
		}
	}
	return facs
}
