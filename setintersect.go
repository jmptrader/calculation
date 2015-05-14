package calc

// Intersect forms the intersection of a variadic number of arguments.
// See https://en.wikipedia.org/wiki/Intersection_%28set_theory%29.
func (setA Set) Intersect(Sets ...Set) Set {
	// Count the number of sets each element appears in.
	cnts := make(map[int]int)
	for elemA := range setA {
		cnts[elemA]++
	}
	for _, setB := range Sets {
		for elemB := range setB {
			cnts[elemB]++
		}
	}

	// The intersection is the set of all elements whose count is equal to the number of sets.
	setC := make(Set)
	setCnt := len(Sets) + 1
	for elemC, cnt := range cnts {
		if cnt == setCnt {
			setC[elemC] = true
		}
	}
	return setC
}
