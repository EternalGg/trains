package smallerNumbersThanCurrent

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

func SmallerNumbersThanCurrent(nums []int) []int {
	qa, numberRemain, qaRemain := make([]int, len(nums)), make(map[int]int), make(map[int]int)

	for _, num := range nums {
		numberRemain[num]++
	}

	numsSort := mergeSort(nums)
	cutRepeat := make([]int, 0)
	last := numsSort[len(numsSort)-1]
	for _, value := range numsSort {
		if last != value {
			cutRepeat = append(cutRepeat, value)
			last = value
		}
	}
	now := 0
	for _, i2 := range cutRepeat {
		qaRemain[i2] = now
		now += numberRemain[i2]
	}

	for i, num := range nums {
		qa[i] = qaRemain[num]
	}
	return qa
}
