package replaceDigits

func ReplaceDigits(s string) string {
	qa := ""
	for i := 0; i < len(s); i += 2 {
		qa += string(s[i])
		if i+1 == len(s)-1 {
			break
		}
		shift := s[i+1] - 48 + s[i]
		qa += string(shift)
	}
	return qa
}
