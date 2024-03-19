package maximumProduct

func MaximumProduct(nums []int) int {
	nums = mergeSort(nums)
	h1, h2, l1, l2, l3 := nums[0], nums[1], nums[len(nums)-1], nums[len(nums)-2], nums[len(nums)-3]
	return max(h1*h2*l1, l1*l2*l3)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
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
