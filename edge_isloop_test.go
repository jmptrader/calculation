package calc

import "testing"

func TestEdgeIsLoop(t *testing.T) {
	// Make some vertices and edges.
	vertex1 := Vertex(1)
	vertex2 := Vertex(2)
	edge1 := Edge{vertex1, vertex1}
	edge2 := Edge{vertex1, vertex2}

	// Check results.
	if !edge1.IsLoop() {
		t.Error("Expected true, got false")
	}
	if edge2.IsLoop() {
		t.Error("Expected false, got true")
	}
}
