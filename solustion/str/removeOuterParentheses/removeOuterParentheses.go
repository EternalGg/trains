package removeOuterParentheses

func RemoveOuterParentheses(s string) string {
	qa, aCount, cutS, start := "", 0, make([]string, 0), 0

	for i, word := range s {
		switch word {
		case 40:
			if aCount == 0 {
				start = i + 1
			}
			aCount++
		case 41:
			if aCount-1 == 0 {
				cutS = append(cutS, s[start:i])
			}
			aCount--
		}
	}
	for _, cut := range cutS {
		qa += cut
	}
	return qa
}
