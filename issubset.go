package calc

// IsSubset tests if A is a subset of B.
func IsSubset(A Set, B Set) bool {
	for a := range A {
		if !B[a] {
			return false
		}
	}
	return true
}
