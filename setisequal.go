package calc

// IsEqual tests if sets are equal.
// See https://en.wikipedia.org/wiki/Equality_%28mathematics%29.
func (setA Set) IsEqual(setB Set) bool {
	// Make sure each element of A is in B.
	for elemA := range setA {
		if !setB[elemA] {
			return false
		}
	}

	// Make sure each element of B is in A.
	for elemB := range setB {
		if !setA[elemB] {
			return false
		}
	}
	return true
}
