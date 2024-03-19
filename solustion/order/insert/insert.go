package insert

import (
	"fmt"
	"sort"
)

func Insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > result[len(result)-1][1] {
			result = append(result, intervals[i])
		} else if intervals[i][0] <= result[len(result)-1][1] {
			if result[len(result)-1][1] < intervals[i][1] {
				result[len(result)-1][1] = intervals[i][1]
			}
		}
	}
	fmt.Println(result)
	return result
}
