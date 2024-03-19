package findColumnWidth

import "strconv"

func findColumnWidth(grid [][]int) []int {
	result := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if len(strconv.Itoa(grid[i][j])) > result[j] {
				result[j] = len(strconv.Itoa(grid[i][j]))
			}
		}
	}
	return result
}
