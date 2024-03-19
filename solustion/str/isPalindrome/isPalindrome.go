package isPalindrome

import (
	"strconv"
)

func IsPalindrome(x int) bool {
	iToStr := strconv.Itoa(x)
	converse := ConverseString(iToStr)
	if converse == iToStr {
		return true
	}
	return false
}

func ConverseString(str string) string {
	result := ""
	for _, letter := range str {
		result = string(letter) + result
	}
	return result
}
