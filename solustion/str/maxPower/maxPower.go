package maxPower

func MaxPower(s string) int {
	lastKey, lastCount, result := uint8(0), 0, 0
	for i := 0; i < len(s); i++ {
		if lastKey == s[i] {
			lastCount++
		} else {
			if lastCount > result {
				result = lastCount
			}
			lastCount = 1
			lastKey = s[i]
		}
	}
	return result
}
