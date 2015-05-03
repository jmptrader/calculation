package calc

import "testing"

func TestPerm(t *testing.T) {
	exp := 720
	act := Perm(10, 3)
	if exp != act {
		t.Error("Expected", exp, "got", act)
	}
}
