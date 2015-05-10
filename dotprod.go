package calc

// DotProd returns the algebraic dot product of two vectors.
// Note: The two vectors should be of equal length.
// See https://en.wikipedia.org/wiki/Dot_product.
func DotProd(vec1 []float64, vec2 []float64) float64 {
	sum := 0.0
	i := 0
	for i < len(vec1) && i < len(vec2) {
		sum += vec1[i] * vec2[i]
		i++
	}
	return sum
}
