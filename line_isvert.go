package calc

// IsVert checks if a line is vertical.
func (l *Line) IsVert() bool {
	return Near(l.Point1.X, l.Point2.X)
}
