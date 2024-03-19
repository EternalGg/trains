package balancedString

func BalancedString(s string) int {
	result, target, countMap := 0, len(s)/4, make(map[int32]int)
	for _, i2 := range s {
		if i2 == 82 || i2 == 81 || i2 == 87 || i2 == 69 {
			countMap[i2]++
		}
	}
	if target > countMap[82] {
		result += target - countMap[82]
	}
	if target > countMap[81] {
		result += target - countMap[81]
	}
	if target > countMap[87] {
		result += target - countMap[87]
	}
	if target > countMap[69] {
		result += target - countMap[69]
	}
	return result
}
