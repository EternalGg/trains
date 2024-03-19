package CheckPermutation

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m1, m2 := make(map[uint8]int), make(map[uint8]int)
	for i := 0; i < len(s1); i++ {
		m1[s1[i]]++
	}
	for i := 0; i < len(s2); i++ {
		m2[s2[i]]++
	}
	for i, i2 := range m1 {
		if m2[i] != i2 {
			return false
		}
	}
	return true
}
