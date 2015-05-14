package calc

import "testing"

func TestSetIsElem(t *testing.T) {
	setA := ToMap([]int{1, 3})

	// Test some elements.
	if !setA.IsElem(1) || !setA.IsElem(3) {
		t.Error("Expected true, got false")
	}

	// Test some nonelements.
	if setA.IsElem(0) || setA.IsElem(2) || setA.IsElem(4) {
		t.Error("Expected false, got true")
	}
}
