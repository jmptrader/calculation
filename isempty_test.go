package calc

import "testing"

func TestIsEmpty(t *testing.T) {
	A := ToMap([]int{})
	B := ToMap([]int{1, 2, 3})

	// Test an empty set.
	if !IsEmpty(A) {
		t.Error("Expected true, got false")
	}

	// Test a nonempty set.
	if IsEmpty(B) {
		t.Error("Expected false, got true")
	}
}
