package lexicalOrder

import "sort"

func lexicalOrder(n int) []int {
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		result = append(result, i)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}
