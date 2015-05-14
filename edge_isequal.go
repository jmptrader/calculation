package calc

// IsEqual tests whether two edges are equal, regardless of vertex order.
// See https://en.wikipedia.org/wiki/Graph_(mathematics).
func (e1 *Edge) IsEqual(e2 *Edge) bool {
	if e1.VertexA == e2.VertexA && e1.VertexB == e2.VertexB {
		return true
	} else if e1.VertexA == e2.VertexB && e1.VertexB == e2.VertexA {
		return true
	}
	return false
}
