package replaceSpaces

func replaceSpaces(S string, length int) string {
	result := ""
	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			result += "%20"
		} else {
			result += string(S[i])
		}
	}
	return result[:length]
}
