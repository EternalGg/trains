package climbingLeaderboard

import (
	"sort"
)

type NodeList struct {
	Val    int32
	Ranked int
	Next   *NodeList
}

func ClimbingLeaderboard(ranked []int32, player []int32) []int32 {
	// Write your code here
	rMap := make(map[int32]bool)
	for i := 0; i < len(ranked); i++ {
		rMap[ranked[i]] = true
	}

	rankList := make([]int, 0)
	for i, _ := range rMap {
		rankList = append(rankList, int(i))
	}

	sort.Slice(rankList, func(i, j int) bool {
		return rankList[i] > rankList[j]
	})
	result := make([]int32, len(player))

	for i := 0; i < len(player); i++ {
		if rMap[player[i]] == false {
			rMap[player[i]] = true
			rankList = append(rankList, int(player[i]))
			sort.Slice(rankList, func(i, j int) bool {
				return rankList[i] > rankList[j]
			})
		}
		for j := 0; j < len(rankList); j++ {
			if player[i] == int32(rankList[j]) {
				result[i] = int32(j + 1)
			}
		}

	}

	return result
}
