package kWeakestRows

import "testing"

func TestKWeakestRows(t *testing.T) {
	matrix := make([][]int, 0)
	matrix = append(matrix, []int{1, 0})
	matrix = append(matrix, []int{1, 0})
	matrix = append(matrix, []int{1, 0})
	matrix = append(matrix, []int{1, 1})
	KWeakestRows(matrix, 4)
}
