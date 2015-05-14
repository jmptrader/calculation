package calc

// ToSet converts a slice of ints to a Set.
func ToSet(s []int) Set {
	setA := make(Set)
	for _, elemA := range s {
		setA[elemA] = true
	}
	return setA
}
