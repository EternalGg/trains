package countBits

import "strconv"

func countBits(n int) []int {
	qa := make([]int, n+1)

	for i := 0; i < n; i++ {
		count := 0
		now := strconv.FormatInt(int64(i), 2)
		for _, letters := range now {
			if letters == 49 {
				count++
			}
		}
		qa = append(qa, count)
	}
	return qa
}
