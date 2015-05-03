package calc

import "testing"

func TestShuffle(t *testing.T) {
	exp := []int{1, 3, 5, 9}
	expSet := ToMap(exp)
	Shuffle(exp)
	actSet := ToMap(exp)
	if !IsEqual(expSet, actSet) {
		t.Error("Expected", expSet, "got", actSet)
	}
}
