package calc

// IsLoop tests whether an arc is a vertex that points to itself.
// See https://en.wikipedia.org/wiki/Graph_(mathematics).
func (a *Arc) IsLoop() bool {
	return a.VertexA == a.VertexB
}
