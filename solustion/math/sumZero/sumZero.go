package sumZero

func sumZero(n int) []int {
	result, half := make([]int, 0), n/2

	for i := 1; i < half; i++ {
		result = append(result, i, -i)
	}
	if n%2 != 0 {
		result = append(result, 0)
	}
	return result
}
