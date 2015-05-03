package calc

import "math"

// Circle type
type Circle struct {
	Radius float64
}

// Area gives area.
func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Diam gives diameter.
func (c *Circle) Diam() float64 {
	return 2.0 * c.Radius
}

// Perim gives perimeter.
func (c *Circle) Perim() float64 {
	return 2.0 * math.Pi * c.Radius
}
