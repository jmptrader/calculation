package calc

// IsHoriz checks if a line is horizontal.
func (l *Line) IsHoriz() bool {
	return Near(l.Point1.Y, l.Point2.Y)
}
