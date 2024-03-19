package arraySign

func arraySign(nums []int) int {
	nav := 0
	for _, num := range nums {
		if num == 0 {
			return 0
		}
		if num < 0 {
			nav++
		}
	}
	if nav%2 == 0 {
		return 1
	} else {
		return 0
	}
}
