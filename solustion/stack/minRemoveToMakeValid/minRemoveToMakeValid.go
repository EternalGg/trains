package minRemoveToMakeValid

func MinRemoveToMakeValid(s string) string {
	b, time := []byte{}, 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ')':
			if time > 0 {
				time--
				b = append(b, ')')
			}
		case '(':
			time++
			b = append(b, '(')
		default:
			b = append(b, s[i])
		}
	}
	result, rt := "", 0
	for i := len(b) - 1; i >= 0; i-- {
		switch b[i] {
		case ')':
			rt++
			result = string(b[i]) + result
		case '(':
			if rt > 0 {
				rt--
				result = string(b[i]) + result
			}
		default:
			result = string(b[i]) + result
		}
	}
	return result
}
