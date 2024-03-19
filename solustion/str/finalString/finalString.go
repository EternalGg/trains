package finalString

func finalString(s string) string {
	str := ""
	for i := 0; i < len(s); i++ {
		if s[i] == 'i' {
			Reverse(str)
		} else {
			str += string(s[i])
		}
	}
	return str
}
func Reverse(s string) string {
	result := ""
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}
