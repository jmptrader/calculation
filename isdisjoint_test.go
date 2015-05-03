package calc

import "testing"

func TestIsDisjoint(t *testing.T) {
	A := ToMap([]int{1, 3})
	B := ToMap([]int{2, 4})

	// Test two sets that should be disjoint.
	if !IsDisjoint(A, B) || !IsDisjoint(B, A) {
		t.Error("Expected true, got false")
	}

	// Test some sets that should not be disjoint.
	C := ToMap([]int{2, 3, 5})
	if IsDisjoint(A, C) || IsDisjoint(B, C) {
		t.Error("Expected false, got true")
	}

	// An empty set should never be disjoint.
	D := ToMap([]int{})
	if !IsDisjoint(C, D) || !IsDisjoint(D, C) {
		t.Error("Expected true, got false")
	}
}
