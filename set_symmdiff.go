package calc

// SymmDiff returns the symmetric difference of A and B (elements in A or B but not both).
// See https://en.wikipedia.org/wiki/Symmetric_difference.
func (setA Set) SymmDiff(setB Set) Set {
	setC := make(Set)
	for elemA := range setA {
		if !setB[elemA] {
			setC[elemA] = true
		}
	}
	for elemB := range setB {
		if !setA[elemB] {
			setC[elemB] = true
		}
	}
	return setC
}
