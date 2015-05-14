package calc

import "math"

// SurfaceArea gives surface area.
func (s *Sphere) SurfaceArea() float64 {
	return 4.0 * math.Pi * s.Radius * s.Radius
}
