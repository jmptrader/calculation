package calc

// IsLoop tests whether an edge is a vertex that points to itself.
func (e *Edge) IsLoop() bool {
	return e.VertexA == e.VertexB
}
