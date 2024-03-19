package divideString

func DivideString(s string, k int, fill byte) []string {
	result := make([]string, 0)

	for i := 0; i < len(s)/k; i++ {
		result = append(result, s[i*k:(i+1)*k])
	}
	cut := len(s) % k
	if cut != 0 {
		now := s[len(s)-cut:]
		for i := 0; i < k-cut; i++ {
			now += string(fill)
		}
		result = append(result, now)

	}
	return result
}
