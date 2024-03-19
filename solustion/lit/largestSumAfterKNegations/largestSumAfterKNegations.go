package largestSumAfterKNegations

import "fmt"

func LargestSumAfterKNegations(nums []int, k int) int {
	nums = mergeSort(nums)
	result, firstZero := 0, len(nums)-1

	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 && k != 0 {
			k--
			nums[i] = -1 * nums[i]
		} else {
			firstZero = i
			break
		}
	}

	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	if k != 0 && k%2 != 0 {
		if firstZero == 0 {
			result -= nums[firstZero] * 2
		} else if firstZero == len(nums)-1 {
			result -= dMin(nums[firstZero], nums[firstZero-1]) * 2
		} else {
			result -= min(nums[firstZero], nums[firstZero-1], nums[firstZero+1]) * 2
		}

	}
	fmt.Println(result)
	return result
}

func dMin(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func min(a, b, c int) int {
	if a >= b && c >= b {
		return b
	} else if b >= c && a >= c {
		return c
	} else {
		return a
	}
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
