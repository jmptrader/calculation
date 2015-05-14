package calc

import "testing"

func TestIntersect(t *testing.T) {
	setA := ToSet([]int{1, 2, 3, 5})
	setB := ToSet([]int{2, 4, 1})
	setC := ToSet([]int{4, 1, 2, 1})
	setD := ToSet([]int{2, 1})
	if !setD.IsEqual(setA.Intersect(setB, setC)) {
		t.Error("Expected true, got false")
	}
}
