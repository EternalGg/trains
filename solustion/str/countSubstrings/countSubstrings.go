package countSubstrings

func countSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if IsPalindrome(s[i : j+1]) {
				count++
			}
		}
	}
	return count
}

func IsPalindrome(str string) bool {

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}
