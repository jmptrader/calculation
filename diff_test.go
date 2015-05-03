package calc

import "testing"

func TestDiff(t *testing.T) {
	A := ToMap([]int{1, 2, 3, 5})
	B := ToMap([]int{2, 4, 1})
	C := ToMap([]int{5, 3})
	if !IsEqual(Diff(A, B), C) {
		t.Error("Expected true, got false")
	}
}
