package numberOfPairs

func NumberOfPairs(nums []int) []int {
	couple, remain, nMap := 0, len(nums), make(map[int]bool)

	for _, num := range nums {
		if nMap[num] {
			couple++
			remain -= 2
		}
		nMap[num] = !nMap[num]
	}
	return []int{couple, remain}
}
