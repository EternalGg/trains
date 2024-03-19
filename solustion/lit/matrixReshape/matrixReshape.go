package matrixReshape

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if r*c > len(mat)*len(mat[0]) {
		return mat
	}
	list := []int{}
	for _, ints := range mat {
		for j := 0; j < len(ints); j++ {
			list = append(list, ints[j])
		}
	}

	result := make([][]int, r)
	last := 0
	for i := 0; i < len(result); i++ {
		for jh := 0; jh < c; jh++ {
			result[i] = append(result[i], list[last])
			last++
		}
	}
	return result
}
