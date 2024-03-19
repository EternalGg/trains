package onesMinusZeros

import "fmt"

func OnesMinusZeros(grid [][]int) [][]int {
	result := make([][]int, len(grid))
	rowMap, colMap := make(map[int]int), make(map[int]int)
	for i := 0; i < len(grid); i++ {
		row := 0
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				row++
			}
		}
		rowMap[i] = row
	}

	for i := 0; i < len(grid[i]); i++ {
		col := 0
		for j := 0; j < len(grid); j++ {
			fmt.Println(i, j)
			if grid[i][j] == 1 {
				col++
			}
		}
		colMap[i] = col
	}

	for i := 0; i < len(grid); i++ {
		result[i] = make([]int, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			result[i][j] = rowMap[i] + colMap[j] - (len(grid) - rowMap[i]) - (len(grid[i]) - colMap[j])
		}
	}
	return result
}
