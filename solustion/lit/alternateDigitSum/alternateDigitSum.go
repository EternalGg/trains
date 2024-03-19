package alternateDigitSum

import "strconv"

func AlternateDigitSum(n int) int {
	str, result := strconv.Itoa(n), 0
	for i := 0; i < len(str); i++ {
		if i%2 == 0 {
			result += int(str[i] - 48)
		} else {
			result -= int(str[i] - 48)
		}
	}
	return result
}
