package calc

import "testing"

func TestIsSubset(t *testing.T) {
	A := ToMap([]int{1, 3})
	B := ToMap([]int{3, 1, 2})

	// A maps can be a subset regardless of order.
	if !IsSubset(A, B) {
		t.Error("Expected true, got false")
	}

	// A map is not a subset if it has a new element.
	A[4] = true
	if IsSubset(A, B) {
		t.Error("Expected false, got true")
	}

	// An empty set is always a subset.
	C := ToMap([]int{})
	if !IsSubset(C, A) || !IsSubset(C, B) || !IsSubset(C, C) {
		t.Error("Expected true, got false")
	}
}
