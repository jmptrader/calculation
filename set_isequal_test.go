package calc

import "testing"

func TestSetIsEqual(t *testing.T) {
	setA := ToSet([]int{1, 2, 3})
	setB := ToSet([]int{3, 1, 2})

	// Two maps should be equal regardless of order.
	if !setA.IsEqual(setB) {
		t.Error("Expected true, got false")
	}

	// Two maps should be not equal if new element.
	setB[4] = true
	if setA.IsEqual(setB) {
		t.Error("Expected false, got true")
	}

	// Two maps should be not equal if one empty.
	setC := ToSet([]int{})
	if setA.IsEqual(setC) {
		t.Error("Expected false, got true")
	}

	// Two maps should be equal if both empty.
	setD := ToSet([]int{})
	if !setC.IsEqual(setD) {
		t.Error("Expected true, got false")
	}
}
