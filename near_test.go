package calc

import "testing"

func TestNear(t *testing.T) {
	if Near(3.0, 3.001) {
		t.Error("Expected false got true")
	}
	if Near(3.0, 2.999) {
		t.Error("Expected false got true")
	}
	if !Near(3.0, 3.0000001) {
		t.Error("Expected true got false")
	}
	if !Near(3.0, 2.9999999) {
		t.Error("Expected true got false")
	}
	if !Near(3.0, 3.0) {
		t.Error("Expected true got false")
	}
}
