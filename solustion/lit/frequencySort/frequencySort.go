package frequencySort

import "fmt"

func FrequencySort(nums []int) []int {
	nMap, ulist := make(map[int]int), make([]int, 0)
	for i := 0; i < len(nums); i++ {
		nMap[nums[i]]++
		if nMap[nums[i]] == 1 {
			ulist = append(ulist, nums[i])
		}
	}
	ulist = mergeSort(ulist)

	sMap, seqList := make(map[int][]int), make([]int, 0)
	for i := 0; i < len(ulist); i++ {
		sMap[nMap[ulist[i]]] = append(sMap[nMap[ulist[i]]], ulist[i])
		if len(sMap[nMap[ulist[i]]]) == 1 {
			seqList = append(seqList, nMap[ulist[i]])
		}
	}
	seqList = mergeSort(seqList)
	result := make([]int, 0)
	for i := 0; i < len(seqList); i++ {
		for j := len(sMap[seqList[i]]) - 1; j >= 0; j-- {
			for k := 0; k < seqList[i]; k++ {
				result = append(result, sMap[seqList[i]][j])
			}
		}
	}
	fmt.Println(result)
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
