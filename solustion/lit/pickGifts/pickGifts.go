package pickGifts

import "math"

func PickGifts(gifts []int, k int) int64 {
	gifts = mergeSort(gifts)
	kList, otherList := make([]int, 0), make([]int, 0)
	if k > len(gifts) {
		kList = gifts
	} else {
		kList, otherList = gifts[len(gifts)-k:], gifts[:len(gifts)-k]
	}

	for i := 0; i < k; i++ {
		kList[len(kList)-1] = int(math.Pow(float64(kList[len(kList)-1]), 0.5))
		kList = mergeSort(kList)
	}
	result := int64(0)
	for _, value := range otherList {
		result += int64(value)
	}
	for _, value := range kList {
		result += int64(value)
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
