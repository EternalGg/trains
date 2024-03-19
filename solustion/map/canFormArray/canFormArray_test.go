package canFormArray

import "testing"

func TestCanFormArray(t *testing.T) {
	list := []int{12}
	matrix := [][]int{}
	matrix = append(matrix, []int{1})
	CanFormArray(list, matrix)
}
