package runeReserve

func RuneReserve(runes []int) int {
	uList, lMap := make([]int, 0), make(map[int]int)
	for i := 0; i < len(runes); i++ {
		if lMap[runes[i]] == 0 {
			uList = append(uList, runes[i])
		}
		lMap[runes[i]]++
	}

	uList = mergeSort(uList)
	max, last := 0, lMap[uList[0]]

	for i := 1; i < len(uList); i++ {

		if uList[i]-1 != uList[i-1] {
			if last > max {
				max = last
			}
			last = lMap[uList[i]]
		} else {
			last += lMap[uList[i]]
		}

	}
	if last > max {
		return last
	} else {
		return max
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
