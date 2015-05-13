package calc

import "testing"

func TestIntersect(t *testing.T) {
	setA := ToMap([]int{1, 2, 3, 5})
	setB := ToMap([]int{2, 4, 1})
	setC := ToMap([]int{4, 1, 2, 1})
	setD := ToMap([]int{2, 1})
	if !setD.IsEqual(Intersect(setA, setB, setC)) {
		t.Error("Expected true, got false")
	}
}
