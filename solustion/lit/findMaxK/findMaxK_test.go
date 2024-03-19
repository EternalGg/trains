package findMaxK

import "testing"

func findMaxK(nums []int) int {
	nMap, max := make(map[int]bool), 0
	for _, num := range nums {
		nMap[num] = true
		anti := 0 - num
		if nMap[anti] && (num > max || anti > max) {
			if num > 0 {
				max = num
			} else {
				max = anti
			}
		}
	}
	if max == 0 {
		return -1
	}
	return max
}
func TestFindMaxK(t *testing.T) {
	nums := []int{-1, 10, 6, 7, -7, 1}
	findMaxK(nums)
}
