package findNumbers

import "strconv"

func findNumbers(nums []int) int {
	count := 0

	for _, num := range nums {
		now := strconv.Itoa(num)
		if len(now)%2 == 0 {
			count++
		}
	}
	return count
}
