package findDigits

import "strconv"

func FindDigits(n int32) int32 {
	// Write your code here
	result := int32(0)
	str := strconv.Itoa(int(n))
	for i := 0; i < len(str); i++ {
		parseInt, _ := strconv.ParseInt(string(str[i]), 10, 64)
		if n%int32(parseInt) == 0 {
			result++
		}
	}
	return result
}
