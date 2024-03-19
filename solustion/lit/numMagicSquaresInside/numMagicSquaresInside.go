package numMagicSquaresInside

func NumMagicSquaresInside(grid [][]int) int {
	if len(grid) < 3 || len(grid[0]) < 3 {
		return 0
	}
	result := 0
	for i := 0; i < len(grid)-3+1; i++ {

		for j := 0; j < len(grid)-3+1; j++ {
			nMap := make(map[int]bool)

			count := grid[i][j+1] + grid[i][j] + grid[i][j+2]
			left := grid[i][j] + grid[i+1][j+1] + grid[i+2][j+2]
			right := grid[i][j+2] + grid[i+1][j+1] + grid[i+2][j]
			if left != count || right != count {
				goto here
			}
			for k := 0; k < 3; k++ {
				if nMap[grid[i][j+k]] == true || nMap[grid[i+1][j+k]] == true || nMap[grid[i+2][j+k]] == true || grid[i][j+k] == grid[i+2][j+k] || grid[i][j+k] == grid[i+1][j+k] || grid[i+1][j+k] == grid[i+2][j+k] || grid[i][j+k] > 9 || grid[i+1][j+k] > 9 || grid[i+2][j+k] > 9 || grid[i][j+k] < 1 || grid[i+1][j+k] < 1 || grid[i+2][j+k] < 1 {
					goto here
				} else {
					nMap[grid[i][j+k]] = true
					nMap[grid[i+1][j+k]] = true
					nMap[grid[i+2][j+k]] = true
				}
			}
			for k := 0; k < 3; k++ {
				line := grid[i][j+1] + grid[i][j] + grid[i][j+2]
				row := grid[i][j] + grid[i+1][j] + grid[i+2][j]
				if line != count || row != count {
					goto here
				}
			}

			result++
		here:
		}

	}
	return result
}
