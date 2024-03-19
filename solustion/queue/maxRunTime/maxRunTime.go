package maxRunTime

import "sort"

func maxRunTime(n int, batteries []int) int64 {
	sort.Ints(batteries)
	sum := 0
	for i := 0; i < len(batteries); i++ {
		sum += batteries[i]
	}
	for i := len(batteries) - 1; i > 0; i-- {
		if batteries[i] <= (sum / n) {
			return int64(len(batteries) - i)
		}
		sum -= batteries[i]
		n--
	}
	return 0
}
