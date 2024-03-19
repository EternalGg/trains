package findReplaceString

func FindReplaceString(s string, indices []int, sources []string, targets []string) string {
	result, key := "", 0
	isMap, itMap := make(map[int]string), make(map[int]string)
	for i := 0; i < len(indices); i++ {
		isMap[indices[i]] = sources[i]
		itMap[indices[i]] = targets[i]
	}
	indices = mergeSort(indices)

	for i := 0; i < len(s); i++ {
		if key == len(indices) {
			result += string(s[i])
		} else {
			if indices[key] == i {
				length := len(isMap[indices[key]])
				if s[indices[key]:length+indices[key]] == isMap[indices[key]] {
					result += itMap[indices[key]]
					i += length - 1
				} else {
					result += string(s[i])
				}
				key++
			} else {
				result += string(s[i])
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
