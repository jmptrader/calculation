package calc

// Intx forms the intersection of a variadic number of arguments.
func Intx(Sets ...Set) Set {
	// Count the number of sets each element appears in.
	counts := make(map[int]int)
	for _, A := range Sets {
		for a := range A {
			counts[a]++
		}
	}

	// The intersection is the set of all elements whose count is equal to the number of sets.
	B := make(Set)
	for a, count := range counts {
		if count == len(Sets) {
			B[a] = true
		}
	}
	return B
}
