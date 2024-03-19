package minCostClimbingStairs

func MinCostClimbingStairs(cost []int) int {
	result, list := 0, make([]int, 0)
	if len(cost)%2 != 0 {
		cost = append(cost, 0)
	}
	for i := 0; i < len(cost); i += 2 {
		now := min(cost[i], cost[i+1])
		list = append(list, now)
	}
	for _, value := range list {
		result += value
	}
	return result
}

func min(i1, i2 int) int {
	if i1 > i2 {
		return i2
	} else {
		return i1
	}
}
