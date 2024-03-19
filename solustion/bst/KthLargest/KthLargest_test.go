package KthLargest

import "testing"

func TestConstructor(t *testing.T) {
	C := Constructor(3, []int{4, 5, 8, 2})
	C.Add(3)
	C.Add(5)
	C.Add(10)
}
