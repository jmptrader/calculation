package calc

// Union forms the union of a variadic number of arguments.
func Union(Sets ...Set) Set {
	B := make(Set)
	for _, A := range Sets {
		for a := range A {
			B[a] = true
		}
	}
	return B
}
