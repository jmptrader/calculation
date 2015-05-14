package calc

// IsEqual tests whether two arcs are equal, in which vertex order is important.
// See https://en.wikipedia.org/wiki/Graph_(mathematics).
func (a1 *Arc) IsEqual(a2 *Arc) bool {
	if a1.VertexA == a2.VertexA && a1.VertexB == a2.VertexB {
		return true
	}
	return false
}
