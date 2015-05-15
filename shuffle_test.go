package calc

import (
	"testing"

	"github.com/BluntSporks/set-theory"
)

func TestShuffle(t *testing.T) {
	exp := []int{1, 3, 5, 9}
	expSet := set.ToSet(exp)
	Shuffle(exp)
	actSet := set.ToSet(exp)
	if !expSet.IsEqual(actSet) {
		t.Error("Expected", expSet, "got", actSet)
	}
}
