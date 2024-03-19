package kWeakestRows

import "sort"

func KWeakestRows(mat [][]int, k int) []int {
	list := make([]int, len(mat))
	length := len(mat[0])
	for i := 0; i < len(mat); i++ {
		count := 0
		for j := 0; j < length; j++ {
			if mat[i][j] == 1 {
				count++
			}
		}
		list[i] = count
	}
	weekMap := make(map[int][]int)
	for i := 0; i < len(list); i++ {
		weekMap[list[i]] = append(weekMap[list[i]], i)
	}

	uList, bMap := make([]int, 0), make(map[int]bool)
	for i := 0; i < len(list); i++ {
		if !bMap[list[i]] {
			uList = append(uList, list[i])
			bMap[list[i]] = true
		}

	}
	sort.Slice(uList, func(i, j int) bool {
		return uList[i] < uList[j]
	})

	result := make([]int, 0)
	for j := 0; j < len(uList); j++ {
		for i := 0; i < len(weekMap[uList[j]]); i++ {
			if k == 0 {
				return result
			} else {
				result = append(result, weekMap[uList[j]][i])
			}
			k--
		}
	}

	return result
}
