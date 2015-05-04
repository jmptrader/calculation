package calc

import "testing"

func TestFac(t *testing.T) {
	exp := map[int]int{2: 2, 3: 1, 5: 1}
	act := Fac(60)
	for fac, cnt := range act {
		if cnt != exp[fac] {
			t.Error("Expected", exp[fac], "got", cnt)
		}
	}
}
