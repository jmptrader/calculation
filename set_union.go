package calc

// Union forms the union of a variadic number of arguments.
// See https://en.wikipedia.org/wiki/Union_%28set_theory%29.
func (setA Set) Union(Sets ...Set) Set {
	setC := make(Set)
	for elemA := range setA {
		setC[elemA] = true
	}
	for _, setB := range Sets {
		for elemB := range setB {
			setC[elemB] = true
		}
	}
	return setC
}
