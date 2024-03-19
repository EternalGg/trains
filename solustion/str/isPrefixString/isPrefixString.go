package isPrefixString

func isPrefixString(s string, words []string) bool {
	time := 0
	for i := 0; i < len(words); i++ {
		if len(s)-time < len(words[i]) {
			return false
		}
		if time == len(s)-1 {
			return true
		}
		for j := 0; j < len(words[i]); j++ {

			if words[i][j] == s[time] {
				time++
				continue
			} else {
				return false
			}
		}
	}
	if time != len(s)-1 {
		return false
	} else {
		return true
	}

}
