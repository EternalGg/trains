package oddString

func OddString(words []string) string {
	diffList := make([]int, len(words[0])-1)
	one := []byte(words[0])
	for i := 1; i < len(one); i++ {
		diffList[i-1] = int(one[i] - one[i-1])
	}
	last, equal, lastLIst := []byte(words[len(words)-1]), true, make([]int, len(words[0])-1)
	for i := 0; i < len(last); i++ {
		lastLIst[i-1] = int(last[i] - last[i-1])
		if diffList[i-1] != int(last[i]-last[i-1]) {
			equal = false
			break
		}
	}
	if equal {
		for i := 1; i < len(words)-1; i++ {
			now := []byte(words[i])
			for j := 0; j < len(now); j++ {
				if diffList[i-1] != int(now[i]-now[i-1]) {
					return words[i]
				}
			}
		}
	} else {
		two := []byte(words[1])
		for i := 0; i < len(two); i++ {
			if diffList[i-1] != int(two[i]-two[i-1]) {
				return words[0]
			}
			if lastLIst[i-1] != int(two[i]-two[i-1]) {
				return words[len(words)-1]
			}
		}
	}
	return ""
}
