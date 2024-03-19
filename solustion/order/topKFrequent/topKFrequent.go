package topKFrequent

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

func topKFrequent(nums []int, k int) []int {
	M, uList := make(map[int]int), make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if M[nums[i]] == 0 {
			uList = append(uList, nums[i])
		}
		M[nums[i]]++
	}

	frequentMap, frequentList := make(map[int][]int), make([]int, 0)
	for i := 0; i < len(uList); i++ {
		if len(frequentMap[M[uList[i]]]) == 0 {
			frequentList = append(frequentList, M[uList[i]])
			list := make([]int, 1)
			list[0] = uList[i]
			frequentMap[M[uList[i]]] = list
		} else {
			frequentMap[M[uList[i]]] = append(frequentMap[M[uList[i]]], uList[i])
		}
	}

	result := make([]int, 0)
	frequentList = mergeSort(frequentList)
	for true {
		for i := len(frequentList) - 1; i >= 0; i-- {
			for j := 0; j < len(frequentMap[frequentList[i]]); j++ {
				k--
				result = append(result, frequentMap[frequentList[i]][j])
				if k == 0 {
					goto here
				}
			}
		}
	}
here:
	return result
}
