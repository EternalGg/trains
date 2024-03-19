package isToeplitzMatrix

import "fmt"

func isToeplitzMatrix(matrix [][]int) bool {
	if len(matrix) < len(matrix[0]) {
		return false
	}
	compare := matrix[0]

	for i := 0; i < len(matrix); i++ {
		ints := matrix[i]
		for j := 0; j < len(compare); j++ {
			if i+j > len(compare)-1 {
				break
			}
			if ints[i+j] != compare[j] {
				fmt.Println(ints[i+j], compare[j])
				return false
			}
		}
	}

	return true
}
