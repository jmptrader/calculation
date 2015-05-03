package calc

import "testing"

func TestUnion(t *testing.T) {
	A := ToMap([]int{1, 2, 3, 5})
	B := ToMap([]int{2, 4, 1})
	C := ToMap([]int{4, 1, 2, 1})
	D := ToMap([]int{2, 1, 5, 3, 4})

	// Test union..
	if !IsEqual(Union(A, B, C), D) {
		t.Error("Expected true, got false")
	}
}
