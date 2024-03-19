package findTheDistanceValue

import "math"

func FindTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	arr2 = mergeSort(arr2)
	lowst, biggest, result := arr2[0], arr2[len(arr2)-1], 0
	for i := 0; i < len(arr1); i++ {
		if math.Abs(float64(arr1[i]-lowst)) > float64(d) && math.Abs(float64(arr1[i]-biggest)) > float64(d) {
			for j := 1; j < len(arr2)-1; j++ {
				if !(math.Abs(float64(arr1[i]-arr2[j])) > float64(d)) {
					goto here
				}
			}
			result++
		} else {
			continue
		}
	here:
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
