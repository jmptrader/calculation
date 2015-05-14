package calc

// IsElem tests if a is an element of A.
// See https://en.wikipedia.org/wiki/Element_%28mathematics%29.
func (setA Set) IsElem(elemA int) bool {
	return setA[elemA]
}
