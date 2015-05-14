package calc

// DotProd returns the algebraic dot product of two vectors.
// Note: The two vectors should be of equal length.
// See https://en.wikipedia.org/wiki/Dot_product.
func (vecA Vec) DotProd(vecB Vec) float64 {
	sum := 0.0
	i := 0
	for i < len(vecA) && i < len(vecB) {
		sum += vecA[i] * vecB[i]
		i++
	}
	return sum
}
