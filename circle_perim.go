package calc

import "math"

// Perim gives the perimeter of the circle.
func (c *Circle) Perim() float64 {
	return 2.0 * math.Pi * c.Radius
}
