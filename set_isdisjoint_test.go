package calc

import "testing"

func TestSetIsDisjoint(t *testing.T) {
	setA := ToSet([]int{1, 3})
	setB := ToSet([]int{2, 4})

	// Test two sets that should be disjoint.
	if !setA.IsDisjoint(setB) || !setB.IsDisjoint(setA) {
		t.Error("Expected true, got false")
	}

	// Test some sets that should not be disjoint.
	setC := ToSet([]int{2, 3, 5})
	if setA.IsDisjoint(setC) || setB.IsDisjoint(setC) {
		t.Error("Expected false, got true")
	}

	// An empty set should never be disjoint.
	setD := ToSet([]int{})
	if !setC.IsDisjoint(setD) || !setD.IsDisjoint(setC) {
		t.Error("Expected true, got false")
	}
}
