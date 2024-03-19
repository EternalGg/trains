package findJudge

import "fmt"

func findJudge(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}
	list, tMap, pMap := make([]int, 0), make(map[int]int), make(map[int]int)
	for i := 0; i < len(trust); i++ {
		tMap[trust[i][1]]++
		if tMap[trust[i][1]] == n-1 {
			list = append(list, trust[i][1])
		}
		pMap[trust[i][0]]++
	}

	for i := 0; i < len(list); i++ {
		fmt.Println(tMap[list[i]], pMap[list[i]])
		if tMap[list[i]] == n-1 && pMap[list[i]] == 0 {
			return list[i]
		}
	}

	return -1

}
