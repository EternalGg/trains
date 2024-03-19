package maxSum

import "testing"

func TestMaxSum(t *testing.T) {
	matrix := make([][]int, 4)
	matrix[0] = []int{6, 2, 1, 3}
	matrix[1] = []int{4, 2, 1, 5}
	matrix[2] = []int{9, 2, 8, 7}
	matrix[3] = []int{4, 1, 2, 9}
	MaxSum(matrix)
}
