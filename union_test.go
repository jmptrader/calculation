package calc

import "testing"

func TestUnion(t *testing.T) {
	setA := ToSet([]int{1, 2, 3, 5})
	setB := ToSet([]int{2, 4, 1})
	setC := ToSet([]int{4, 1, 2, 1})
	setD := ToSet([]int{2, 1, 5, 3, 4})

	// Test union..
	if !setD.IsEqual(Union(setA, setB, setC)) {
		t.Error("Expected true, got false")
	}
}
