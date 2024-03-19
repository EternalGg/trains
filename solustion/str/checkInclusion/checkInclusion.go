package checkInclusion

// s2 == s  s1 == target
func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	strMap, length, left, right, checkMap, list := make(map[uint8]int), len(s2), 0, 0, make(map[uint8]bool), make([]uint8, 0)
	for _, value := range s1 {
		if !checkMap[uint8(value)] {
			checkMap[uint8(value)] = true
			list = append(list, uint8(value))
		}
		strMap[uint8(value)]++
	}
	tempoMap := make(map[uint8]int)
	for right = 0; right < len(s1); right++ {
		if checkMap[s2[right]] {
			tempoMap[s2[right]]++
		}
	}
	if checkListInMap(strMap, tempoMap, list) {
		return true
	}

	for left = 0; right < length; right++ {
		if checkMap[s2[left]] {
			tempoMap[s2[left]]--
		}
		if checkMap[s2[right]] {
			tempoMap[s2[right]]++
		}

		if checkMap[s2[right]] && checkListInMap(strMap, tempoMap, list) {
			return true
		}
		left++
	}
	return false
}

func checkListInMap(m, c map[uint8]int, nums []uint8) bool {
	for _, checkValue := range nums {
		if c[checkValue] != m[checkValue] {
			return false
		}
	}
	return true
}
