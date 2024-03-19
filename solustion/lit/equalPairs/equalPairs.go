package equalPairs

import (
	"fmt"
	"strconv"
)

func equalPairs(grid [][]int) int {
	rowMap := make(map[string]bool)
	for i := 0; i < len(grid); i++ {
		str := ""
		for j := 0; j < len(grid[i]); j++ {
			str += strconv.Itoa(grid[i][j]) + ","
		}
		fmt.Println(str, i)
		rowMap[str] = true
	}
	result := 0
	for i := 0; i < len(grid[0]); i++ {
		str := ""
		for j := 0; j < len(grid); j++ {
			str += strconv.Itoa(grid[i][j]) + ","
		}

		if rowMap[str] {
			fmt.Println(str)
			result++
		}
	}
	return result
}
