package findWinners

import "sort"

func FindWinners(matches [][]int) [][]int {
	result, cMap, uMap, l := make([][]int, 2), make(map[int]int), make(map[int]bool), make([]int, 0)
	for i := 0; i < len(matches); i++ {
		if uMap[matches[i][0]] == false {
			uMap[matches[i][0]] = true
			l = append(l, matches[i][0])
		}
		if uMap[matches[i][1]] == false {
			uMap[matches[i][1]] = true
			l = append(l, matches[i][1])
		}
		if cMap[matches[i][0]] == 0 {
			cMap[matches[i][0]] = -1
		}
		if cMap[matches[i][1]] == -1 {
			cMap[matches[i][1]] = 1
		} else {
			cMap[matches[i][1]]++
		}
	}
	//delete(uMap)
	sort.Slice(l, func(i, j int) bool {
		return l[i] < l[j]
	})
	for _, i := range l {
		if cMap[i] == -1 {
			result[0] = append(result[0], i)
		}
		if cMap[i] == 1 {
			result[1] = append(result[1], i)
		}
	}
	return result
}
