package calc

import "testing"

func TestStdDevPop(t *testing.T) {
	exp := 2.141845
	act := StdDevPop([]float64{2.3, 3.5, 5.7, 7.9})
	if !Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
