package smallestEvenMultiple

import "testing"

func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	} else {
		return 2 * n
	}
}
func TestSmallestEvenMultiple(t *testing.T) {
	smallestEvenMultiple(6)
}
