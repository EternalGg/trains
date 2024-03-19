package firstPalindrome

import "fmt"

func firstPalindrome(words []string) string {
	for i := 0; i < len(words); i++ {
		if IsPalindrome(words[i]) {
			return words[i]
		}
	}
	return ""
}

func IsPalindrome(s string) bool {
	sToJustLetter := stringJustLetter(s)
	converse := converseString(sToJustLetter)
	fmt.Println(sToJustLetter, converse)
	if converse == sToJustLetter {
		return true
	}
	return false
}

func stringJustLetter(str string) string {
	result := ""
	for _, letter := range str {
		if letter >= 65 && letter <= 90 {
			result += string(letter + 32)
			continue
		}
		if (letter >= 97 && letter <= 122) || (letter >= 48 && letter <= 57) {
			result += string(letter)
			continue
		}
	}
	return result
}

func converseString(str string) string {
	result := ""
	for _, letter := range str {
		result = string(letter) + result
	}
	return result
}
