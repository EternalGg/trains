package maxSum

func MaxSum(grid [][]int) int {
	max, answer := 0, make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		answer[i] = make([]int, len(grid[i])-2)
		for j := 0; j < len(grid[i])-2; j++ {
			answer[i][j] = grid[i][j] + grid[i][j+1] + grid[i][j+2]
		}
	}
	for i := 0; i < len(answer)-2; i++ {

		for j := 0; j < len(answer[i]); j++ {
			now := 0
			now += answer[i][j]
			now += answer[i+2][j]
			now += grid[i+1][j+1]
			if now > max {
				max = now
			}
		}

	}
	return max
}
