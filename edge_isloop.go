package calc

// IsLoop tests whether an edge is a vertex that points to itself.
// See https://en.wikipedia.org/wiki/Graph_(mathematics).
func (e *Edge) IsLoop() bool {
	return e.VertexA == e.VertexB
}
