package checkString

func CheckString(s string) bool {
	b := false
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' && b != false {
			return false
		}
		if s[i] == 'b' {
			b = true
		}

	}
	return true
}
