package maxIncreaseKeepingSkyline

func maxIncreaseKeepingSkyline(grid [][]int) int {
	rowMax, lawMax := make(map[int]int), make(map[int]int)
	for i := 0; i < len(grid); i++ {
		row := mergeSort(grid[i])
		rowMax[i] = row[len(row)-1]
		law := make([]int, len(grid))
		for j := 0; j < len(grid); j++ {
			law = append(law, grid[j][i])
		}
		law = mergeSort(law)
		lawMax[i] = law[len(law)-1]
	}

	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			law, row, min := lawMax[j], rowMax[i], 0
			if law < row {
				min = law
			} else {
				min = row
			}
			if grid[i][j] < min {
				count += min - grid[i][j]
			}
		}
	}
	return count
}

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		// 分治，两两拆分，一直拆到基础元素才向上递归。
		return nums
	}
	i := len(nums) / 2
	left := mergeSort(nums[0:i])
	// 左侧数据递归拆分
	right := mergeSort(nums[i:])
	// 右侧数据递归拆分
	result := merge(left, right)
	// 排序 & 合并
	return result
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	i, j := 0, 0
	l, r := len(left), len(right)
	for i < l && j < r {
		if left[i] > right[j] {
			result = append(result, right[j])
			j++
		} else {
			result = append(result, left[i])
			i++
		}
	}
	result = append(result, right[j:]...)
	result = append(result, left[i:]...)
	return result
}
