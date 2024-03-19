package findLucky

func findLucky(arr []int) int {
	unqiueList, countMap := make([]int, 0), make(map[int]int)
	arr = mergeSort(arr)
	for _, value := range arr {
		if countMap[value] == 0 {
			unqiueList = append(unqiueList, value)
		}
		countMap[value]++
	}
	for i := len(unqiueList) - 1; i > 0; i-- {
		if unqiueList[i] == countMap[unqiueList[i]] {
			return unqiueList[i]
		}
	}
	return -1
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
