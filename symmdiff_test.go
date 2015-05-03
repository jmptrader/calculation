package calc

import "testing"

func TestSymmDiff(t *testing.T) {
	A := ToMap([]int{1, 2, 3, 5})
	B := ToMap([]int{2, 4, 1})
	C := ToMap([]int{5, 4, 3})
	if !IsEqual(SymmDiff(A, B), C) {
		t.Error("Expected true, got false")
	}
}
