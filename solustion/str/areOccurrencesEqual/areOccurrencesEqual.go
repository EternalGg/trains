package areOccurrencesEqual

func AreOccurrencesEqual(s string) bool {
	i32Map := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		i32Map[s[i]]++
	}
	avg := i32Map[(s[0])]
	for i := 0; i < len(s); i++ {
		if i32Map[s[i]] != avg {
			return false
		}
	}
	return true
}
