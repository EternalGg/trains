package validPalindrome

func ValidPalindrome(s string) bool {
	sMap, list := make(map[int32]int), make([]int32, 0)

	for _, i2 := range s {
		if sMap[i2] == 0 {
			list = append(list, i2)
		}
		sMap[i2]++
	}
	odd := 0
	for _, i2 := range list {
		if sMap[i2]%2 != 0 {
			odd++
		}
		if odd > 2 {
			return false
		}
	}
	return true
}
