package calc

import "testing"

func TestFac(t *testing.T) {
	exp := map[int]int{2: 2, 3: 1, 5: 1}
	act := Fac(60)
	for fac, count := range act {
		if count != exp[fac] {
			t.Error("Expected", exp[fac], "got", count)
		}
	}
}
