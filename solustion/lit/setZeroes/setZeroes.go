package setZeroes

func setZeroes(matrix [][]int) {
	raw, row := make(map[int]bool), make(map[int]bool)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				raw[i] = true
				row[j] = true
			}
		}
	}
	for i, _ := range raw {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = 0
		}
	}
	for i, _ := range row {
		for j := 0; j < len(matrix); j++ {
			matrix[j][i] = 0
		}
	}
}
