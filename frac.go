package calc

type Frac struct {
	Numer int
	Denom int
}

// Add sums two fractions.
func Add(frac1 Frac, frac2 Frac) Frac {
	// Adjust fractions to common denominator.
	numer := frac1.Numer*frac2.Denom + frac1.Denom*frac2.Numer
	denom := frac1.Denom * frac2.Denom

	// Yield result.
	return Frac{numer, denom}

}

// Sub subtracts two fractions.
func Sub(frac1 Frac, frac2 Frac) Frac {
	// Adjust fractions to common denominator.
	numer := frac1.Numer*frac2.Denom - frac1.Denom*frac2.Numer
	denom := frac1.Denom * frac2.Denom

	// Yield result.
	return Frac{numer, denom}

}

// Mul multiplies two fractions.
func Mul(frac1 Frac, frac2 Frac) Frac {
	// Adjust fractions to common denominator.
	numer := frac1.Numer * frac2.Numer
	denom := frac1.Denom * frac2.Denom

	// Yield result.
	return Frac{numer, denom}

}

// Div divides two fractions.
func Div(frac1 Frac, frac2 Frac) Frac {
	// Adjust fractions to common denominator.
	numer := frac1.Numer * frac2.Denom
	denom := frac1.Denom * frac2.Numer

	// Yield result.
	return Frac{numer, denom}

}
