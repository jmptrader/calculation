package calc

// IsDisjoint tests if A is disjoint from (has no common elements with) B.
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
