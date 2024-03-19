package findLeastNumOfUniqueInts

import (
	"sort"
)

func FindLeastNumOfUniqueInts(arr []int, k int) int {
	cMap, list := make(map[int]int), make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if cMap[arr[i]] == 0 {
			list = append(list, arr[i])
		}
		cMap[arr[i]]++
	}
	count := len(list)
	nMap, nList := make(map[int][]int), make([]int, 0)
	for i := 0; i < len(list); i++ {
		if len(nMap[cMap[list[i]]]) == 0 {
			nList = append(nList, cMap[list[i]])
			nMap[cMap[list[i]]] = make([]int, 0)
		}
		nMap[cMap[list[i]]] = append(nMap[cMap[list[i]]], list[i])
	}
	sort.Slice(nList, func(i, j int) bool {
		return nList[i] < nList[j]
	})

	for i := 0; i < len(nList); i++ {
		for j := 0; j < len(nMap[nList[i]]); j++ {
			k -= nList[i]
			if k == 0 {
				count--
				return count
			} else if k < 0 {
				return count
			} else if k > 0 {
				count--
			}
		}
	}
	return count
}
