package calc

import "math"

// Sphere type
type Sphere struct {
	Radius float64
}

// Diam gives diameter.
func (s *Sphere) Diam() float64 {
	return 2.0 * s.Radius
}

// SurfArea gives surface area.
func (s *Sphere) SurfArea() float64 {
	return 4.0 * math.Pi * s.Radius * s.Radius
}

// Vol gives volume.
func (s *Sphere) Vol() float64 {
	return 4.0 / 3.0 * math.Pi * s.Radius * s.Radius * s.Radius
}
