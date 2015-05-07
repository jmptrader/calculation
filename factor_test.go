package calc

import "testing"

func TestFactor(t *testing.T) {
	exp := map[int]int{2: 2, 3: 1, 5: 1}
	act := Factor(60)
	for factor, cnt := range act {
		if cnt != exp[factor] {
			t.Error("Expected", exp[factor], "got", cnt)
		}
	}
}
