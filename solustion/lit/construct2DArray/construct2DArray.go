package construct2DArray

func Construct2DArray(original []int, m int, n int) [][]int {
	if n*m > len(original) {
		return [][]int{}
	}
	result, key, remain := make([][]int, m), 0, 0
	for i, _ := range result {
		for j := 0; j < n; j++ {
			if key == len(original) {
				goto here
			}
			result[i] = append(result[i], original[key])
			key++
		}
		remain = i + 1
	}
here:
	return result[:remain]
}
