package calc

// Diff returns the relative complement, or difference, of sets A and B.
func Diff(A Set, B Set) Set {
	C := make(Set)
	for a := range A {
		if !B[a] {
			C[a] = true
		}
	}
	return C
}
