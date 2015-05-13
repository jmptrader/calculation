package calc

import "math"

const (
	b0 = 0.2316419
	b1 = 0.319381530
	b2 = -0.356563782
	b3 = 1.781477937
	b4 = -1.821255978
	b5 = 1.330274429
)

// CFI returns Phi, the cumulative density function of the PDF.
// See https://en.wikipedia.org/wiki/Normal_distribution.
// The return value of this function can be used more flexibly than the empirical rule.
// To convert CDF to empirical value, use 1 - (1-CDF(x))*2.
// See https://en.wikipedia.org/wiki/68%E2%80%9395%E2%80%9399.7_rule.
func CDF(x float64) float64 {
	x = math.Abs(x)
	t := 1.0 / (1.0 + b0*x)
	tPow := t
	factor := b1 * tPow
	tPow *= t
	factor += b2 * tPow
	tPow *= t
	factor += b3 * tPow
	tPow *= t
	factor += b4 * tPow
	tPow *= t
	factor += b5 * tPow
	return 1.0 - PDF(x)*factor
}
