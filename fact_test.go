package calc

import "testing"

func TestFact(t *testing.T) {
	exp := 1
	facts := Fact()
	for i, fact := range facts {
		if i > 0 {
			exp *= i
		}
		if exp != fact {
			t.Error("Expected", exp, "got", fact)
		}
	}
}
