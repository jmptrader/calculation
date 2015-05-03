package calc

import "math"

// QuadRoot calculates the real roots of a quadratic equation.
func QuadRoot(a float64, b float64, c float64) []float64 {
	roots := make([]float64, 0, 2)

	// Calculate the discriminant.
	disc := b*b - 4.0*a*c
	if disc < 0.0 {
		return roots
	}
	r := math.Sqrt(disc)

	// Calculate the first root.
	x1 := (-1*b + r) / (2.0 * a)
	roots = append(roots, x1)
	if Near(disc, 0.0) {
		return roots
	}

	// Calculate the second root
	x2 := (-1*b - r) / (2.0 * a)
	roots = append(roots, x2)
	return roots
}
