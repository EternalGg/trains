package shiftGrid

import "train/solustion/lit/construct2DArray"

func shiftGrid(grid [][]int, k int) [][]int {
	count := len(grid) * len(grid[0])
	last := k % count
	if k == count || k == 0 || last == 0 {
		return grid
	}
	list := make([]int, 0)
	for _, ints := range grid {
		for j := 0; j < len(ints); j++ {
			list = append(list, ints[j])
		}
	}

	clist := make([]int, 0)
	for i := count - last; i < count; i++ {
		clist = append(clist, list[i])
	}
	for i := 0; i < count-last; i++ {
		clist = append(clist, list[i])
	}
	result := construct2DArray.Construct2DArray(clist, len(grid), len(grid[0]))
	return result
}
