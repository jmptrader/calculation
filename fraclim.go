package calc

import (
	"math"
	"sort"
)

// FracLim returns the limiting value below which a fraction of the list lies.
// If fraction is 1.0, the maximum is returned.
// If fraction is 0.5, the median is returned.
// If fraction is 0.0, the minimum is returned.
func FracLim(values []float64, frac float64) float64 {
	// Sort the list so we can split it.
	sort.Float64s(values)

	// Declare result.
	var fracLim float64

	// Get the index location.
	n := len(values)
	indexFloat := frac * float64(n-1)
	indexInt := int(indexFloat)

	// See if we are very close to an int, and if so, round, otherwise, split.
	isInt := math.Abs(float64(indexInt)-indexFloat) < 0.001
	if isInt {
		fracLim = values[indexInt]
	} else {
		floor := int(math.Floor(indexFloat))
		ceil := int(math.Ceil(indexFloat))
		fracLim = (values[floor] + values[ceil]) / 2.0
	}
	return fracLim
}
