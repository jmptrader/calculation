package calc

import "testing"

func TestFacl(t *testing.T) {
	exp := 1
	facls := Facl()
	for i, facl := range facls {
		if i > 0 {
			exp *= i
		}
		if exp != facl {
			t.Error("Expected", exp, "got", facl)
		}
	}
}
