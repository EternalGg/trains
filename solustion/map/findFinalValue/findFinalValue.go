package findFinalValue

func FindFinalValue(nums []int, original int) int {
	nMap := make(map[int]bool)
	for _, num := range nums {
		nMap[num] = true
	}
	for nMap[original] {
		original *= 2
	}
	return original / 2
}
