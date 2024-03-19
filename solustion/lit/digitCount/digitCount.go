package digitCount

func digitCount(num string) bool {
	nMap := make(map[int]int)
	for _, i2 := range num {
		nMap[int(i2-48)]++
	}
	for i, i2 := range num {
		if nMap[i] != int(i2)-48 {
			return false
		}
	}
	return true
}
