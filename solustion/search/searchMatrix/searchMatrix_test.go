package searchMatrix

import "testing"

func TestBinary(t *testing.T) {
	list := make([][]int, 3)
	list[0] = []int{1}
	list[1] = []int{3}
	list[2] = []int{5}
	SearchMatrix(list, 3)
}
