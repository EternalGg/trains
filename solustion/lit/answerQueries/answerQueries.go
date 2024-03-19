package answerQueries

import "fmt"

func answerQueries(nums []int, queries []int) []int {
	nums = mergeSort(nums)
	cList := make([]int, len(nums))
	cList[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		cList[i] += cList[i-1]
	}
	fmt.Println(cList)
	result := make([]int, 0)
	for _, query := range queries {
		result = append(result, BinaryFindSmall(cList, query))
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

func BinaryFindSmall(nums []int, target int) int {
	result := BinaryFind(nums, 0, len(nums)-1, target)
	return result
}

func BinaryFind(nums []int, head, last, target int) int {
	half := ((last - head) / 2) + head
	if head > last {
		return half - 1
	}
	switch {
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
	return half
}
