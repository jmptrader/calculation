package calc

// IsSubset tests if A is a subset of B.
// See https://en.wikipedia.org/wiki/Subset.
func (setA Set) IsSubset(setB Set) bool {
	for elemA := range setA {
		if !setB[elemA] {
			return false
		}
	}
	return true
}
