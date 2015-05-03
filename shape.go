package calc

// Shape is the interface of a 2-dimensional enclosed figure.
type Shape interface {
	Area() float64
	Perim() float64
}
