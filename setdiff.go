package calc

// Diff returns the relative complement, or difference, of sets A and B.
// See https://en.wikipedia.org/wiki/Complement_%28set_theory%29.
func (setA Set) Diff(setB Set) Set {
	setC := make(Set)
	for elemA := range setA {
		if !setB[elemA] {
			setC[elemA] = true
		}
	}
	return setC
}
