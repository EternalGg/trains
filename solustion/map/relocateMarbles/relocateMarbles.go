package relocateMarbles

import (
	"fmt"
	"sort"
)

func relocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for i := 0; i < len(moveFrom); i++ {
		if m[moveFrom[i]] > 0 {
			c := m[moveFrom[i]]
			m[moveTo[i]] += c
			if m[moveTo[i]] == 0 {
				delete(m, moveFrom[i])
			}
		}
	}
	result := []int{}
	for key, value := range m {
		fmt.Println(value, key)
		result = append(result, key)
	}
	sort.Ints(result)
	return result
}
