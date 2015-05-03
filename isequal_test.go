package calc

import "testing"

func TestIsEqual(t *testing.T) {
	A := ToMap([]int{1, 2, 3})
	B := ToMap([]int{3, 1, 2})

	// Two maps should be equal regardless of order.
	if !IsEqual(A, B) {
		t.Error("Expected true, got false")
	}

	// Two maps should be not equal if new element.
	B[4] = true
	if IsEqual(A, B) {
		t.Error("Expected false, got true")
	}

	// Two maps should be not equal if one empty.
	C := ToMap([]int{})
	if IsEqual(A, C) {
		t.Error("Expected false, got true")
	}

	// Two maps should be equal if both empty.
	D := ToMap([]int{})
	if !IsEqual(C, D) {
		t.Error("Expected true, got false")
	}
}
