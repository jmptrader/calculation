package calc

import "testing"

func TestCircle(t *testing.T) {
	expArea := 28.27433388
	expDiam := 6.0
	expPerim := 18.84955592

	c := Circle{3.0}
	actArea := c.Area()
	actDiam := c.Diam()
	actPerim := c.Perim()

	if !Near(expArea, actArea) {
		t.Error("Expected", expArea, "got", actArea)
	}
	if !Near(expDiam, actDiam) {
		t.Error("Expected", expDiam, "got", actDiam)
	}
	if !Near(expPerim, actPerim) {
		t.Error("Expected", expPerim, "got", actPerim)
	}
}
