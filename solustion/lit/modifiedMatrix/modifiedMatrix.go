package modifiedMatrix

func modifiedMatrix(matrix [][]int) [][]int {
	maxMap := map[int]int{}
	for i := 0; i < len(matrix[0]); i++ {
		max := -1
		for j := 0; j < len(matrix); j++ {
			if matrix[j][i] > max {
				max = matrix[j][i]
			}
		}
		maxMap[i] = max
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == -1 {
				matrix[i][j] = maxMap[j]
			}
		}
	}
	return matrix
}
