package calc

// Intersect forms the intersection of a variadic number of arguments.
// See https://en.wikipedia.org/wiki/Intersection_%28set_theory%29.
func Intersect(Sets ...Set) Set {
	// Count the number of sets each element appears in.
	cnts := make(map[int]int)
	for _, A := range Sets {
		for a := range A {
			cnts[a]++
		}
	}

	// The intersection is the set of all elements whose count is equal to the number of sets.
	B := make(Set)
	for a, cnt := range cnts {
		if cnt == len(Sets) {
			B[a] = true
		}
	}
	return B
}
