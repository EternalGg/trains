package minimizedStringLength

func minimizedStringLength(s string) int {
	m := make(map[int32]bool)
	result := 0
	for _, i2 := range s {
		if !m[i2] {
			result++
			m[i2] = true
		}
	}
	return result
}
