package calc

// Div divides two fractions.
func Div(frac1 Frac, frac2 Frac) Frac {
	// Adjust fractions to common denominator.
	num := frac1.Num * frac2.Den
	den := frac1.Den * frac2.Num

	// Yield result.
	return Frac{num, den}

}
