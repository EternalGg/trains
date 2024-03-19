package countNegatives

func countNegatives(grid [][]int) int {
	result := 0
	for i := 0; i < len(grid); i++ {
		if grid[i][len(grid[i])-1] < 0 {
			result += FindFirstNegative(grid[i])
		}
	}
	return result
}

func FindFirstNegative(list []int) int {
	if list[0] < 0 {
		return len(list)
	}
	max, min := len(list)-1, 1
	target := 0
	for !((list[min] < 0 && list[min-1] >= 0) || (list[max] < 0 && list[max-1] >= 0)) {
		target = ((max - min) / 2) + min
		if list[target] >= 0 {
			min = target
		} else {
			max = target
		}
	}
	if list[min] < 0 && list[min-1] >= 0 {
		return len(list) - min
	} else {
		return len(list) - max
	}

}
