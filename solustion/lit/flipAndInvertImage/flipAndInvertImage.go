package flipAndInvertImage

func flipAndInvertImage(image [][]int) [][]int {
	result := make([][]int, len(image))
	for i := 0; i < len(image); i++ {
		for j := len(image[i]) - 1; j >= 0; j-- {
			if image[i][j] == 0 {
				result[i] = append(result[i], 1)
			} else {
				result[i] = append(result[i], 0)
			}
		}
	}
	return result
}
