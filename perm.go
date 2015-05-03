package calc

// Perm returns the permuation count of r objects taken from n.
func Perm(n int, r int) int {
	perm := 1
	for i := n - r + 1; i <= n; i++ {
		perm *= i
	}
	return perm
}
