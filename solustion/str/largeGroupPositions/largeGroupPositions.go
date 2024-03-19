package largeGroupPositions

func LargeGroupPositions(s string) [][]int {
	last, time, result := s[0], 0, make([][]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == last {
			time++
		} else {
			if time >= 3 {
				result = append(result, []int{i - time, i - 1})
			}
			last = s[i]
			time = 1
		}
	}
	if time >= 3 {
		result = append(result, []int{len(s) - time, len(s) - 1})
	}

	return result
}
