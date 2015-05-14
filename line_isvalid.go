package calc

// IsValid checks if the line is a valid line definition of two different points.
func (l *Line) IsValid() bool {
	return !Near(l.Point1.X, l.Point2.X) || !Near(l.Point1.Y, l.Point2.Y)
}
