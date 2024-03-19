package arrayRankTransform

func ArrayRankTransform(arr []int) []int {
	rd := removeDuplicate(arr)
	rd = mergeSort(rd)
	index := make(map[int]int)
	for i, i2 := range rd {
		index[i2] = i
	}
	result := make([]int, len(arr))
	for i, i2 := range arr {
		result[i] = index[i2]
	}
	return result
}

func removeDuplicate(list []int) []int {
	unique, result := make(map[int]bool), make([]int, 0)
	for _, i2 := range list {
		if unique[i2] == false {
			result = append(result, i2)
			unique[i2] = true
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
