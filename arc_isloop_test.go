package calc

import "testing"

func TestArcIsLoop(t *testing.T) {
	// Make some vertices and arcs.
	vertex1 := Vertex(1)
	vertex2 := Vertex(2)
	arc1 := Arc{vertex1, vertex1}
	arc2 := Arc{vertex1, vertex2}

	// Check results.
	if !arc1.IsLoop() {
		t.Error("Expected true, got false")
	}
	if arc2.IsLoop() {
		t.Error("Expected false, got true")
	}
}
