package calc

// IsDisjoint tests if A is disjoint from (has no common elements with) B.
// See https://en.wikipedia.org/wiki/Disjoint_sets.
func (setA Set) IsDisjoint(setB Set) bool {
	for elemA := range setA {
		if setB[elemA] {
			return false
		}
	}
	for elemB := range setB {
		if setA[elemB] {
			return false
		}
	}
	return true
}
