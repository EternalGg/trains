package mergeSimilarItems

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	list, lMap := make([]int, 0), make(map[int]int)

	for i := 0; i < len(items1); i++ {
		list = append(list, items1[i][0])
		lMap[items1[i][0]] = items1[i][1]
	}

	for i := 0; i < len(items2); i++ {
		if lMap[items1[i][0]] == 0 {
			list = append(list, items2[i][0])
			lMap[items2[i][0]] = items2[i][1]
		} else {
			lMap[items2[i][0]] += items2[i][1]
		}
	}
	result := make([][]int, 0)
	list = mergeSort(list)
	for i := 0; i < len(list); i++ {
		now := make([]int, 0)
		now = append(now, list[i], lMap[list[i]])
		result = append(result, now)
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
