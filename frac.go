package calc

// Fraction type.
// See https://en.wikipedia.org/wiki/Fraction_%28mathematics%29.
type Frac struct {
	Num int
	Den int
}

// Add sums two fractions.
func Add(frac1 Frac, frac2 Frac) Frac {
	// Adjust fractions to common denominator.
	num := frac1.Num*frac2.Den + frac1.Den*frac2.Num
	den := frac1.Den * frac2.Den

	// Yield result.
	return Frac{num, den}

}

// Sub subtracts two fractions.
func Sub(frac1 Frac, frac2 Frac) Frac {
	// Adjust fractions to common denominator.
	num := frac1.Num*frac2.Den - frac1.Den*frac2.Num
	den := frac1.Den * frac2.Den

	// Yield result.
	return Frac{num, den}

}

// Mul multiplies two fractions.
func Mul(frac1 Frac, frac2 Frac) Frac {
	// Adjust fractions to common denominator.
	num := frac1.Num * frac2.Num
	den := frac1.Den * frac2.Den

	// Yield result.
	return Frac{num, den}

}

// Div divides two fractions.
func Div(frac1 Frac, frac2 Frac) Frac {
	// Adjust fractions to common denominator.
	num := frac1.Num * frac2.Den
	den := frac1.Den * frac2.Num

	// Yield result.
	return Frac{num, den}

}
