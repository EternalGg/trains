package oddCells

func oddCells(m int, n int, indices [][]int) int {
	matrix, qa := [][]int{}, 0
	for i := 0; i < m; i++ {
		newRaw := make([]int, n)
		matrix = append(matrix, newRaw)
	}

	for _, index := range indices {
		for i := 0; i < n; i++ {
			matrix[index[0]][i] += 1
		}
		for i := 0; i < m; i++ {
			matrix[i][index[1]] += 1
		}
	}

	for _, nums := range matrix {
		for _, num := range nums {
			if num%2 != 0 {
				qa++
			}
		}
	}
	return qa
}
