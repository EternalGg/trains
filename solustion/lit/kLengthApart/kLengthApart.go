package kLengthApart

func kLengthApart(nums []int, k int) bool {
	lastKey := -1
	for key, num := range nums {
		if num == 1 {
			if lastKey == -1 {
				lastKey = key
				continue
			} else {
				if key-lastKey < k {
					return false
				}
				lastKey = key
			}
		}
	}
	return true
}
