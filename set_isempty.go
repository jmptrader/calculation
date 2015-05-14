package calc

// IsEmpty tests if set A is empty.
// See https://en.wikipedia.org/wiki/Empty_set.
func (setA Set) IsEmpty() bool {
	return len(setA) == 0
}
