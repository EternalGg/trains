package largestOddNumber

func LargestOddNumber(num string) string {
	if num[len(num)-1]%2 != 0 {
		return num
	}
	last, length := 0, 0
	for i := 0; i < len(num); i++ {
		if num[i]%2 != 0 {
			if length == 0 {
				length++
			} else {
				last = i
			}
		}
	}
	if length == 0 {
		return ""
	}
	return num[:last+1]
}
