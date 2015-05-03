package calc

import "testing"

func TestIntx(t *testing.T) {
	A := ToMap([]int{1, 2, 3, 5})
	B := ToMap([]int{2, 4, 1})
	C := ToMap([]int{4, 1, 2, 1})
	D := ToMap([]int{2, 1})
	if !IsEqual(Intx(A, B, C), D) {
		t.Error("Expected true, got false")
	}
}
