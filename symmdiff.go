package calc

// SymmDiff returns the symmetric difference of A and B (elements in A or B but not both).
func SymmDiff(A Set, B Set) Set {
	C := make(Set)
	for a := range A {
		if !B[a] {
			C[a] = true
		}
	}
	for b := range B {
		if !A[b] {
			C[b] = true
		}
	}
	return C
}
