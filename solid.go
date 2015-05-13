package calc

// Solid is the interface of a 3-dimensional enclosed figure.
// See https://en.wikipedia.org/wiki/Three-dimensional_space.
type Solid interface {
	SurfaceArea() float64
	Vol() float64
}
