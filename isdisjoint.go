package calc

// IsDisjoint tests if A is disjoint from (has no common elements with) B.
// See https://en.wikipedia.org/wiki/Disjoint_sets.
func IsDisjoint(A Set, B Set) bool {
	for a := range A {
		if B[a] {
			return false
		}
	}
	for b := range B {
		if A[b] {
			return false
		}
	}
	return true
}
