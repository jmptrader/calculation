package calc

import "testing"

func TestArcIsEqual(t *testing.T) {
	// Make some vertices and arcs.
	vertex1 := Vertex(1)
	vertex2 := Vertex(2)
	arc1 := Arc{vertex1, vertex2}
	arc2 := Arc{vertex1, vertex2}
	arc3 := Arc{vertex2, vertex1}

	// Check results.
	if !arc1.IsEqual(&arc2) {
		t.Error("Expected true, got false")
	}
	if !arc2.IsEqual(&arc1) {
		t.Error("Expected true, got false")
	}
	if !arc1.IsEqual(&arc1) {
		t.Error("Expected true, got false")
	}
	if arc1.IsEqual(&arc3) {
		t.Error("Expected false, got true")
	}
}
