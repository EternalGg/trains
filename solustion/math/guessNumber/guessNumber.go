package guessNumber

import "fmt"

func guess(num int) int {
	// 我的比较小
	if num < 6 {
		return -1
	}
	// 我的比较大
	if num > 6 {
		return 1
	}

	// 一样
	if num == 6 {
		return 0
	}
	return 0
}

// -1 : target > now 比目标小
// 1 :  target < now 比目标大
func GuessNumber(n int) int {

	max, min, target := n, 1, n
	for guess(target) != 0 {
		target = ((max - min) / 2) + min
		if guess(target) == -1 {
			min = target
		} else {
			max = target
		}
	}
	fmt.Println(target)
	return target
}
