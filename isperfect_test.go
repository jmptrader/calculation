package calc

import "testing"

func TestIsPerfect(t *testing.T) {
	for i := 2; i < 100; i++ {
		exp := (i == 6 || i == 28)
		act := IsPerfect(i)
		if exp != act {
			t.Error("Expected", exp, "got", act)
		}
	}
}
