package cutTheSticks

func cutTheSticks(arr []int32) []int32 {
	// Write your code here
	max := int32(0)
	lenght := len(arr)
	m := make(map[int32]int)
	count := int32(0)
	for i := 0; i < lenght; i++ {
		if arr[i] >= max {
			max = arr[i]
		}
		m[arr[i]]++
	}
	result := make([]int32, 0)
	for i := 0; i < int(max); i++ {
		result = append(result, int32(lenght)-count)
		if result[i] == 1 {
			break
		}
		count += int32(m[int32(i+1)])
	}
	return result
}
