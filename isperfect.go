package calc

// IsPerfect checks if an integer is a perfect number, equal to the sum of its proper divisors.
// See https://en.wikipedia.org/wiki/Perfect_number.
func IsPerfect(num int) bool {
	sum := 0
	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum += i
		}
	}
	return sum == num
}
