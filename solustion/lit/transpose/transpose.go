package transpose

func Transpose(matrix [][]int) [][]int {
	result := make([][]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			result[j] = append(result[j], matrix[i][j])
		}
	}
	return result
}
