package maxScore

func MaxScore(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 49 {
			count++
		}
	}
	max := 0
	for i := 0; i < len(s); i++ {
		if s[i] != 49 {
			count++
		} else {
			count--
		}
		if count > max {
			max = count
		}
	}
	return max
}
