package calc

// IsEqual tests if sets are equal.
func IsEqual(A Set, B Set) bool {
	for a := range A {
		if !B[a] {
			return false
		}
	}
	for b := range B {
		if !A[b] {
			return false
		}
	}
	return true
}
