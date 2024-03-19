package topKFrequent

import "sort"

func TopKFrequent(words []string, k int) []string {
	wMap, uList := make(map[string]int), []string{}
	for i := 0; i < len(words); i++ {
		wMap[words[i]]++
		if wMap[words[i]] == 1 {
			uList = append(uList, words[i])
		}
	}

	cMap, tMap, cList := make(map[int][]string), make(map[int]bool), make([]int, 0)
	for i := 0; i < len(uList); i++ {
		if len(cMap[wMap[uList[i]]]) == 0 {
			cMap[wMap[uList[i]]] = make([]string, 0)
		}
		if !tMap[wMap[uList[i]]] {
			cList = append(cList, wMap[uList[i]])
			tMap[wMap[uList[i]]] = true
		}
		cMap[wMap[uList[i]]] = append(cMap[wMap[uList[i]]], uList[i])
	}

	sort.Slice(cList, func(i, j int) bool {
		return cList[i] > cList[j]
	})
	result := []string{}
	for i := 0; i < len(cList); i++ {
		if k != 0 {
			if k >= len(cMap[cList[i]]) {
				sort.Slice(cMap[cList[i]], func(j, k int) bool {
					return cMap[cList[i]][j] < cMap[cList[i]][k]
				})
				for j := 0; j < len(cMap[cList[i]]); j++ {
					result = append(result, cMap[cList[i]][j])
				}
				k -= len(cMap[cList[i]])
			} else {
				sort.Slice(cMap[cList[i]], func(j, k int) bool {
					return cMap[cList[i]][j] < cMap[cList[i]][k]
				})
				for j := 0; j < k; j++ {
					result = append(result, cMap[cList[i]][j])
				}
				break
			}
		}
	}

	return result
}
