package calc

import "testing"

func TestSetDiff(t *testing.T) {
	setA := ToSet([]int{1, 2, 3, 5})
	setB := ToSet([]int{2, 4, 1})
	setC := ToSet([]int{5, 3})
	if !setC.IsEqual(setA.Diff(setB)) {
		t.Error("Expected true, got false")
	}
}
