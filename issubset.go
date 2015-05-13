package calc

// IsSubset tests if A is a subset of B.
// See https://en.wikipedia.org/wiki/Subset.
func IsSubset(A Set, B Set) bool {
	for a := range A {
		if !B[a] {
			return false
		}
	}
	return true
}
