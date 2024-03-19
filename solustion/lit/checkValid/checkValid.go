package checkValid

import "fmt"

func CheckValid(matrix [][]int) bool {
	for i := 0; i < len(matrix); i++ {
		rawMap, rawCount := make(map[int]bool), len(matrix[i])
		for j := 0; j < len(matrix[i]); j++ {
			if !rawMap[matrix[i][j]] {
				rawCount--
				rawMap[matrix[i][j]] = true
			}
		}
		fmt.Println(rawCount)
		if rawCount != 0 {
			return false
		}

		lawMap, lawCount := make(map[int]bool), len(matrix[i])
		for j := 0; j < len(matrix[i]); j++ {
			if !lawMap[matrix[j][i]] {
				rawCount--
				lawMap[matrix[j][i]] = true
			}
		}
		fmt.Println(lawMap)
		if lawCount != 0 {
			return false
		}
	}
	return true
}
