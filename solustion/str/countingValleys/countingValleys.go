package countingValleys

func countingValleys(steps int32, path string) int32 {
	// Write your code here
	isDown, result := false, 0
	if path[0] == 'D' {
		isDown = true
	}
	for i := 0; i < int(steps); i++ {
		if isDown && path[i] == 'U' {
			result++
		}
		if path[i] == 'D' {
			isDown = true
		}
	}
	return int32(result)
}
func countingValleys1(steps int32, path string) int32 {
	result, count := int32(0), 0
	if path[0] == 'D' {
		count--
	} else {
		count++
	}
	for i := 1; i < len(path); i++ {
		if path[i] == 'U' && count+1 == 0 {
			count++
			result++
			continue
		}
		if path[i] == 'D' {
			count--
			continue
		}
		if path[i] == 'U' {
			count++
			continue
		}
	}
	return result
}
