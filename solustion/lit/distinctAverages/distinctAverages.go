package distinctAverages

func DistinctAverages(nums []int) int {
	//unMap, newList := make(map[int]bool), make([]int, 0)
	//for i := 0; i < len(nums); i++ {
	//	if unMap[nums[i]] == false {
	//		newList = append(newList, nums[i])
	//		unMap[nums[i]] = true
	//	}
	//}
	nums = mergeSort(nums)

	uMap, result := make(map[float32]bool), 0

	for i := 0; i < len(nums); i++ {
		last := len(nums) - i - 1
		if last <= i {
			break
		}
		now := float32(nums[last]+nums[i]) / float32(2)
		if uMap[now] == false {
			uMap[now] = true
			result++
		}
	}
	return result
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
