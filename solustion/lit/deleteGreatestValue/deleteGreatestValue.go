package deleteGreatestValue

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

func DeleteGreatestValue(grid [][]int) int {
	result, length := 0, len(grid[0])
	for i := 0; i < len(grid); i++ {
		grid[i] = mergeSort(grid[i])
	}
	for i := 0; i < length; i++ {
		max := 0
		for _, ints := range grid {
			if ints[i] > max {
				max = ints[i]
			}
		}
		result += max
	}
	return result
}
