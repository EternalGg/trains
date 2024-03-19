package countDigits

import "strconv"

func CountDigits(num int) int {
	str, result := strconv.Itoa(num), 0
	sToByte := []byte(str)

	for _, b := range sToByte {
		now := int(b) - 48
		if num%now == 0 {
			result++
		}
	}
	return result
}
