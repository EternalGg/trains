package repeatedSubstringPattern

func RepeatedSubstringPattern(s string) bool {
	length, half := len(s), len(s)/2
	for i := half; i > 0; i-- {
		if length%i == 0 {
			now := s[:i]
			for j := 0; j < length/i; j++ {
				if s[i*j:i*(j+1)] != now {
					goto here
				}
			}
			return true
		here:
		}
	}
	return false
}
