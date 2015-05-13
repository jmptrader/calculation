package calc

// IsEmpty tests if set A is empty.
// See https://en.wikipedia.org/wiki/Empty_set.
func IsEmpty(A Set) bool {
	return len(A) == 0
}
