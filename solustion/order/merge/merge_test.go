package merge

import "testing"

func TestMerge(t *testing.T) {
	matrix := [][]int{}
	matrix = append(matrix, []int{4, 5})
	matrix = append(matrix, []int{1, 4})

	matrix = append(matrix, []int{8, 10})

	matrix = append(matrix, []int{15, 18})
	Merge(matrix)
}
