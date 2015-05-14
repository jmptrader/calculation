package calc

import "math"

// Area gives area.
// See http://www.mathopenref.com/coordtrianglearea.html.
func (t *Triangle) Area() float64 {
	ax := t.Point1.X
	ay := t.Point1.Y
	bx := t.Point2.X
	by := t.Point2.Y
	cx := t.Point3.X
	cy := t.Point3.Y
	return math.Abs(ax*(by-cy)+bx*(cy-ay)+cx*(ay-by)) / 2.0
}
