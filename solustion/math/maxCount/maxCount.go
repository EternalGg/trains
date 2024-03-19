package maxCount

func MaxCount(banned []int, n int, maxSum int) int {
	result, count, point, M := 0, 0, 0, make(map[int]bool)
	for _, value := range banned {
		M[value] = true
	}
	for point != n {
		point++
		if M[point] == true {
			continue
		} else {
			if count+point > maxSum {
				return result
			} else {
				count += point
				result++
			}
		}
	}
	return result
}
