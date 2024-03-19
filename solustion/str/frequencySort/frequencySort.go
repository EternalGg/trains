package frequencySort

func FrequencySort(s string) string {
	bMap, bList := make(map[byte]int), make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if bMap[s[i]] == 0 {
			bList = append(bList, s[i])
		}
		bMap[s[i]]++
	}

	sortList, sortMap := make([]int, 0), make(map[int][]byte)
	for i := 0; i < len(bList); i++ {
		if len(sortMap[bMap[bList[i]]]) == 0 {
			blist := make([]byte, 0)
			sortMap[bMap[bList[i]]] = blist
			sortList = append(sortList, bMap[bList[i]])
		}
		sortMap[bMap[bList[i]]] = append(sortMap[bMap[bList[i]]], bList[i])
	}

	sortList = mergeSort(sortList)

	result := ""
	for i := len(sortList) - 1; i >= 0; i-- {
		// len(list)次
		for j := 0; j < len(sortMap[sortList[i]]); j++ {
			// len（map[list]次)
			for k := 0; k < sortList[i]; k++ {
				result += string(sortMap[sortList[i]][j])
			}
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
