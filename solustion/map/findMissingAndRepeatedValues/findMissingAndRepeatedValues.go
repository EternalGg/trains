package findMissingAndRepeatedValues

func FindMissingAndRepeatedValues(grid [][]int) []int {
	result := []int{0, 0}
	list := make([]int, len(grid)*len(grid)+1)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			list[grid[i][j]]++
		}
	}
	for i := 0; i < len(list); i++ {
		if list[i] != 1 {
			if list[i] == 2 {
				result[0] = i
			}
			if list[i] == 0 {
				result[1] = i
			}
		}
	}
	return result
}
