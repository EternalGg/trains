package shiftingLetters

func ShiftingLetters(s string, shifts []int) string {

	count := 0
	for i := 0; i < len(shifts); i++ {
		count += shifts[i]
	}
	result, sToB := "", []byte(s)
	for i := 0; i < len(sToB); i++ {
		now := count % 26
		answer := sToB[i] + byte(now)
		if answer > byte(122) {
			answer = answer%122 + 96
		}
		result += string(answer)
		count -= shifts[i]
	}
	return result
}
