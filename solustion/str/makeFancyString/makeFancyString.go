package makeFancyString

func MakeFancyString(s string) string {
	result, lastValue, lastCount := "", uint8(0), 0

	for i := 0; i < len(s); i++ {
		if s[i] == lastValue {
			lastCount++
		} else {
			if lastCount > 2 {
				lastCount = 2
			}
			for j := 0; j < lastCount; j++ {
				result += string(lastValue)
			}
			lastValue = s[i]
			lastCount = 1
		}
	}

	if lastCount > 2 {
		lastCount = 2
	}
	for j := 0; j < lastCount; j++ {
		result += string(lastValue)
	}

	return result
}
