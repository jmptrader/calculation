package calc

// Add sums two fractions.
func Add(frac1 Frac, frac2 Frac) Frac {
	// Adjust fractions to common denominator.
	num := frac1.Num*frac2.Den + frac1.Den*frac2.Num
	den := frac1.Den * frac2.Den

	// Yield result.
	return Frac{num, den}

}
