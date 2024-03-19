package mostFrequentEven

func mostFrequentEven(nums []int) int {
	evenList, evenMap, evenFrequen := make([]int, 0), make(map[int]bool), make(map[int]int)

	for _, num := range nums {
		if num%2 == 0 {
			if !evenMap[num] {
				evenList = append(evenList, num)
				evenMap[num] = true
			}
			evenFrequen[num]++
		}
	}
	evenList = mergeSort(evenList)
	if len(evenList) == 0 {
		return -1
	}
	max, maxList := 1, make([]int, 0)
	for _, value := range evenList {
		if evenFrequen[value] > max {
			max = evenFrequen[value]
			maxList = make([]int, 0)
			continue
		}
		if evenFrequen[value] == max {
			maxList = append(maxList, value)
			continue
		}
	}

	return maxList[0]
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
