package calc

// IsPerp checks if l1 is perpendicular to l2.
func (l1 *Line) IsPerp(l2 *Line) bool {
	if l1.IsVert() && l2.IsHoriz() {
		return true
	}
	if l1.IsHoriz() && l2.IsVert() {
		return true
	}
	return Near(l1.Slope()*l2.Slope(), -1.0)
}
