package maxIceCream

import "sort"

func maxIceCream(costs []int, coins int) int {
	sort.Slice(costs, func(i, j int) bool {
		return costs[i] < costs[j]
	})
	result := 0
	for i := 0; i < len(costs); i++ {
		if coins-costs[i] >= 0 {
			result++
			coins -= costs[i]
		}
	}
	return result
}
