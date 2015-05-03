package calc

import "testing"

func TestFibonacci(t *testing.T) {
	exp := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	seq := Fibonacci(12)
	for i, sum := range seq {
		if exp[i] != sum {
			t.Error("Expected", exp[i], "got", sum)
		}
	}
}
