package buildArray

func buildArray(target []int, n int) []string {
	last, result := 1, make([]string, 0)
	push, pop := "push", "pop"
	for _, i2 := range target {

		if i2 > last {
			for i := 0; i < i2-last; i++ {
				result = append(result, push)
				result = append(result, pop)
			}
		} else if i2 == last {
			result = append(result, push)
		}

		last = i2
	}
	return result
}
