package calc

import "testing"

func TestSetDiff(t *testing.T) {
	setA := ToMap([]int{1, 2, 3, 5})
	setB := ToMap([]int{2, 4, 1})
	setC := ToMap([]int{5, 3})
	if !IsEqual(setA.Diff(setB), setC) {
		t.Error("Expected true, got false")
	}
}
