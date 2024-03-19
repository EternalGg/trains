package findNthDigit

import "strconv"

func findNthDigit(n int) int {
	str, now := "", 1
	for len(str) != n {
		str += strconv.Itoa(now)
		now++
	}
	result, _ := strconv.ParseInt(string(str[n]), 10, 64)
	return int(result)
}
