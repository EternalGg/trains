package minimumAbsDifference

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

func minimumAbsDifference(arr []int) [][]int {
	result := [][]int{}
	arr = mergeSort(arr)
	min := arr[1] - arr[0]
	for i := 1; i < len(arr); i++ {
		cut := arr[i] - arr[i-1]
		if cut > min {
			continue
		}
		if cut < min {
			min = cut
			result = [][]int{}
		}
		now := []int{arr[i-1], arr[i]}
		result = append(result, now)
	}
	return result
}
