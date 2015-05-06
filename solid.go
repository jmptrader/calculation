package calc

// Solid is the interface of a 3-dimensional enclosed figure.
type Solid interface {
	SurfaceArea() float64
	Vol() float64
}
