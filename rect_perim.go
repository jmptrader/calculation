package calc

// Perim gives perimeter.
func (r *Rect) Perim() float64 {
	return 2.0*r.Length + 2.0*r.Width
}
