package calc

import "testing"

func TestSetIsEmpty(t *testing.T) {
	setA := ToSet([]int{})
	setB := ToSet([]int{1, 2, 3})

	// Test an empty set.
	if !setA.IsEmpty() {
		t.Error("Expected true, got false")
	}

	// Test a nonempty set.
	if setB.IsEmpty() {
		t.Error("Expected false, got true")
	}
}
