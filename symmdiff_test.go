package calc

import "testing"

func TestSymmDiff(t *testing.T) {
	setA := ToMap([]int{1, 2, 3, 5})
	setB := ToMap([]int{2, 4, 1})
	setC := ToMap([]int{5, 4, 3})
	if !setC.IsEqual(SymmDiff(setA, setB)) {
		t.Error("Expected true, got false")
	}
}
