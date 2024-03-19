package findRepeatNumber

func findRepeatNumber(nums []int) int {
	M := make(map[int]bool)
	for _, num := range nums {
		if M[num] == true {
			return num
		} else {
			M[num] = true
		}
	}
	return 0
}
