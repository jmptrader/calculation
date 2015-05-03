package calc

import "testing"

func TestStirling(t *testing.T) {
	exp := 1.0
	for i := 2; i < 20; i++ {
		act := Stirling(i)
		exp *= float64(i)
		if act < 0.95*exp || act > 1.05*exp {
			t.Error("Expected", exp, "got", act)
		}
	}
}
