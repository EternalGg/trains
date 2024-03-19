package maxEnvelopes

import (
	"fmt"
	"sort"
)

func MaxEnvelopes(envelopes [][]int) int {
	//l1, uMap := make([]int, 0), make(map[int][]int)
	//for i := 0; i < len(envelopes); i++ {
	//	if len(uMap[envelopes[i][0]]) == 0 {
	//		l1 = append(l1, envelopes[i][0])
	//		uMap[envelopes[i][0]] = make([]int, 0)
	//	}
	//	uMap[envelopes[i][0]] = append(uMap[envelopes[i][0]], envelopes[i][1])
	//}
	//l1 = mergeSort(l1)
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})

	fmt.Println(envelopes)
	return 1
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

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})

	f := []int{}
	for _, e := range envelopes {
		h := e[1]
		if i := sort.SearchInts(f, h); i < len(f) {
			f[i] = h
		} else {
			f = append(f, h)
		}
	}
	return len(f)
}
