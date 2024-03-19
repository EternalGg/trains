package birthday

func Birthday(s []int32, d int32, m int32) int32 {
	// Write your code here
	result := 0
	for i := 0; i < len(s); i++ {
		count := int32(0)
		for j := i; j < len(s); j++ {
			count += s[i]
			if count == d && j-i+1 == int(m) {
				result++
			}
			if count > d {
				break
			}
		}
	}
	return int32(result)
}
