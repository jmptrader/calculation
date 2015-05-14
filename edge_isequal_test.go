package calc

import "testing"

func TestEdgeIsEqual(t *testing.T) {
	// Make some vertices and edges.
	vertex1 := Vertex(1)
	vertex2 := Vertex(2)
	vertex3 := Vertex(3)
	edge1 := Edge{vertex1, vertex2}
	edge2 := Edge{vertex2, vertex1}
	edge3 := Edge{vertex1, vertex3}

	// Check results.
	if !edge1.IsEqual(&edge2) {
		t.Error("Expected true, got false")
	}
	if !edge2.IsEqual(&edge1) {
		t.Error("Expected true, got false")
	}
	if !edge1.IsEqual(&edge1) {
		t.Error("Expected true, got false")
	}
	if edge1.IsEqual(&edge3) {
		t.Error("Expected false, got true")
	}
}
