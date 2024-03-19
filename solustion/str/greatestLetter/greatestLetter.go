package greatestLetter

func GreatestLetter(s string) string {
	sMap, result := make(map[uint8]bool), uint8(0)

	for i := 0; i < len(s); i++ {
		if s[i] >= 64 && s[i] <= 90 {
			if sMap[s[i]+32] == true && result < s[i] {
				result = s[i]
			}
			sMap[s[i]] = true
		} else if s[i] >= 97 && s[i] <= 122 {
			if sMap[s[i]-32] == true && result < s[i]-32 {
				result = s[i] - 32
			}
			sMap[s[i]] = true
		}
	}
	if result == 0 {
		return ""
	}
	return string(result)
}
