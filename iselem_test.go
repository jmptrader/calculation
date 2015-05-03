package calc

import "testing"

func TestIsElem(t *testing.T) {
	A := ToMap([]int{1, 3})

	// Test some elements.
	if !IsElem(1, A) || !IsElem(3, A) {
		t.Error("Expected true, got false")
	}

	// Test some nonelements.
	if IsElem(0, A) || IsElem(2, A) || IsElem(4, A) {
		t.Error("Expected false, got true")
	}
}
