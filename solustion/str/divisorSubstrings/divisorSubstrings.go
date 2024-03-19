package divisorSubstrings

import "strconv"

func DivisorSubstrings(num int, k int) int {
	itoa := strconv.Itoa(num)
	result := 0
	for i := 0; i < len(itoa)-k+1; i++ {
		pi, _ := strconv.ParseInt(itoa[i:i+k], 10, 64)
		if pi == 0 {
			continue
		}
		if int64(num)%pi == 0 {
			result++
		}
	}
	return result
}
