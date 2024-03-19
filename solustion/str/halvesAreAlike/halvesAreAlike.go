package halvesAreAlike

func HalvesAreAlike(s string) bool {
	A, B := 0, 0
	half := len(s) / 2
	for i := 0; i < half; i++ {
		switch s[i] {
		case 65, 69, 73, 79, 85, 97, 101, 105, 111, 117:
			A++
		}
	}
	for i := half; i < len(s); i++ {
		switch s[i] {
		case 65, 69, 73, 79, 85, 97, 101, 105, 111, 117:
			B++
		}
	}
	if A == B {
		return true
	} else {
		return false
	}
}
