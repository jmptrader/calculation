package calc

import "testing"

func TestRect(t *testing.T) {
	expArea := 6.0
	expPerim := 10.0

	r := Rect{2.0, 3.0}
	actArea := r.Area()
	actPerim := r.Perim()

	if !Near(expArea, actArea) {
		t.Error("Expected", expArea, "got", actArea)
	}
	if !Near(expPerim, actPerim) {
		t.Error("Expected", expPerim, "got", actPerim)
	}
}
