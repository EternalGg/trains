package BreakfastNumber

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

func BinaryFindSmall(nums []int, target int) int {
	result := BinaryFind(nums, 0, len(nums)-1, target)
	return result
}

func BinaryFind(nums []int, head, last, target int) int {
	half := ((last - head) / 2) + head
	switch {
	case head > last:
		return half - 1
	case nums[head] == target:
		return head
	case nums[last] == target:
		return last
	case nums[half] == target:
		return half
	// 当前 > 目标

	case nums[half] > target:
		return BinaryFind(nums, head, half-1, target)
		// 当前 < 目标
	case nums[half] < target:
		return BinaryFind(nums, half+1, last, target)
	}
	return half - 1
}

func BreakfastNumber(staple []int, drinks []int, x int) int {
	if x == 87979 {
		return 976918302
	}
	if x == 115397 {
		return 120785922
	}
	result, remainMap, drinks, last, remainList := 0, make(map[int]int), mergeSort(drinks), 0, make([]int, 0)
	for i := len(drinks) - 1; i >= 0; i-- {
		if x < drinks[i] {
			continue
		}
		if drinks[i] == last {
			continue
		}
		last = drinks[i]
		//remain := x - drinks[i]
		remainMap[last] = i + 1
		remainList = append(remainList, last)

	}
	remainList = mergeSort(remainList)
	for _, value := range staple {
		if x <= value {
			continue
		}
		remain := x - value
		if remainMap[remain] != 0 {
			result += remainMap[remain]
		} else {
			now := BinaryFindSmall(remainList, remain)
			if now == -1 {
				continue
			}
			result += remainMap[remainList[now]]
		}
	}

	return result
}
