package checkValid

func checkValid(matrix [][]int) bool {
	uMap := make(map[int]int)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] >= len(matrix[i]) || matrix[i][j] <= 0 {
				return false
			} else {
				uMap[matrix[i][j]]++
				if uMap[matrix[i][j]] != i {
					return false
				}
			}
		}
	}
	return true
}
