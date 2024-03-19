package removeTrailingZeros

func removeTrailingZeros(num string) string {
	length := len(num) - 1
	for i := len(num) - 1; i >= 0; i-- {
		if num[i] == '0' {
			length--
		} else {
			break
		}
	}
	return num[:length]
}
