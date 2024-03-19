package rowAndMaximumOnes

func rowAndMaximumOnes(mat [][]int) []int {
	max, maxKey := 0, 0
	for i := 0; i < len(mat); i++ {
		now := 0
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				now++
			}
		}
		if now > max {
			maxKey = i
		}
	}

	return []int{maxKey, max}
}
