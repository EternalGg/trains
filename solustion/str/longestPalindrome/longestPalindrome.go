package longestPalindrome

func LongestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	result := ""
	for i := len(s); i > 0; i-- {
		for j := 0; j < len(s); j++ {
			if j+i-1 >= len(s) {
				break
			}
			for IsPalindrome(s[j : i+j]) {
				result = s[j : i+j]
				return result
			}
		}
	}
	return result
}

func IsPalindrome(str string) bool {

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}
