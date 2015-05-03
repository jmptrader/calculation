package calc

// ToMap converts a slice of ints to a map to bool.
func ToMap(s []int) Set {
	A := make(Set)
	for _, a := range s {
		A[a] = true
	}
	return A
}
