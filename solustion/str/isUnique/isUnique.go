package isUnique

func isUnique(astr string) bool {
	m := make(map[uint8]bool)
	for i := 0; i < len(astr); i++ {
		if m[astr[i]] == true {
			return false
		}
	}
	return true
}
