package calc

// IsParallel checks if l1 is parallel to l2.
func (l1 *Line) IsParallel(l2 *Line) bool {
	if l1.IsVert() && l2.IsVert() {
		return true
	}
	return Near(l1.Slope(), l2.Slope())
}
