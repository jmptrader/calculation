package calc

import "math"

// Vol gives volume.
func (s *Sphere) Vol() float64 {
	return 4.0 / 3.0 * math.Pi * s.Radius * s.Radius * s.Radius
}
