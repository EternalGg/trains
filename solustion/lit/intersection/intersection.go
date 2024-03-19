package intersection

func Intersection(nums [][]int) []int {
	M := make(map[int]bool)
	for _, i2 := range nums[0] {
		M[i2] = true
	}
	for i := 1; i < len(nums); i++ {
		M = DeathLoop(M, nums[i])
	}
	result := make([]int, 0)
	for i := 0; i < len(nums[0]); i++ {
		if M[nums[0][i]] {
			result = append(result, nums[0][i])
		}
	}
	result = mergeSort(result)
	return result
}

func DeathLoop(m map[int]bool, nums []int) map[int]bool {
	newMap := make(map[int]bool)
	for _, num := range nums {
		if m[num] {
			newMap[num] = true
		}
	}
	return newMap
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
