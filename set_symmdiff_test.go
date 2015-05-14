package calc

import "testing"

func TestSymmDiff(t *testing.T) {
	setA := ToSet([]int{1, 2, 3, 5})
	setB := ToSet([]int{2, 4, 1})
	setC := ToSet([]int{5, 4, 3})
	if !setC.IsEqual(setA.SymmDiff(setB)) {
		t.Error("Expected true, got false")
	}
}
