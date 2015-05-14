package calc

import "math"

// Area gives the area of the circle.
func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
