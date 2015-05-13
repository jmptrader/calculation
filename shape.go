package calc

// Shape is the interface of a 2-dimensional enclosed figure.
// See https://en.wikipedia.org/wiki/Shape.
type Shape interface {
	Area() float64
	Perim() float64
}
