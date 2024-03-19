package numberOfSteps

func NumberOfSteps(num int) int {
	count := 0

	for num != 1 {
		if num%2 == 0 {
			num = num / 2
			count++
		} else {
			num -= 1
			count++
		}
	}

	return count
}
