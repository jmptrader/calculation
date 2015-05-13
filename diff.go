package calc

// Diff returns the relative complement, or difference, of sets A and B.
// See https://en.wikipedia.org/wiki/Complement_%28set_theory%29.
func Diff(A Set, B Set) Set {
	C := make(Set)
	for a := range A {
		if !B[a] {
			C[a] = true
		}
	}
	return C
}
