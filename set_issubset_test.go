package calc

import "testing"

func TestIsSubset(t *testing.T) {
	setA := ToMap([]int{1, 3})
	setB := ToMap([]int{3, 1, 2})

	// A map can be a subset regardless of order.
	if !setA.IsSubset(setB) {
		t.Error("Expected true, got false")
	}

	// A map is not a subset if it has a new element.
	setA[4] = true
	if setA.IsSubset(setB) {
		t.Error("Expected false, got true")
	}

	// An empty set is always a subset.
	setC := ToMap([]int{})
	if !setC.IsSubset(setA) || !setC.IsSubset(setB) || !setC.IsSubset(setC) {
		t.Error("Expected true, got false")
	}
}
