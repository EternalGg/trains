package isStraight

func isStraight(nums []int) bool {
	nums = mergeSort(nums)
	zero := 0
	for _, num := range nums {
		if num == 0 {
			zero++
		}
	}
	max, min := nums[4], nums[4]
	for i := 4; i >= 0; i-- {
		if nums[i] < min {
			min = nums[i]
		}
	}
	switch zero {
	case 0:
		for i := 0; i < len(nums)-1; i++ {
			if nums[i]+1 != nums[i+1] {
				return false
			}
		}
	case 1:

	case 2:

	case 3:
		if max-min != 4 {
			return false
		}
	default:
		return true
	}
	return true
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
