package calc

// Fibonacci returns the Fibonacci sequence to the nth term.
// See https://en.wikipedia.org/wiki/Fibonacci_number.
func Fibonacci(n int) []int {
	seq := make([]int, n)
	seq[0] = 0
	seq[1] = 1
	for i := 2; i < n; i++ {
		seq[i] = seq[i-2] + seq[i-1]
	}
	return seq
}
