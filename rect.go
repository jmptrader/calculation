package calc

// Rectangle type
type Rect struct {
	Length, Width float64
}

// Area gives area.
func (r *Rect) Area() float64 {
	return r.Length * r.Width
}

// Perim gives perimeter.
func (r *Rect) Perim() float64 {
	return 2.0*r.Length + 2.0*r.Width
}
