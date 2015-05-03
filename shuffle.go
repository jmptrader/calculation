package calc

import (
	"math/rand"
	"time"
)

// Shuffle a list of numbers in place.
// See https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle for details of modern method.
func Shuffle(array []int) {
	rand.Seed(time.Now().UnixNano())
	for i := len(array) - 1; i > 0; i-- {
		j := rand.Intn(i)
		array[i], array[j] = array[j], array[i]
	}
}
