package repeatedNTimes

func RepeatedNTimes(nums []int) int {
	target, repeatMap := len(nums)/2, make(map[int]int)
	for _, num := range nums {
		repeatMap[num]++
	}
	for _, num := range nums {
		if target == repeatMap[num] {
			return num
		}
	}
	return 0
}
