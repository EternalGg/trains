package equalizeArray

func equalizeArray(arr []int32) int32 {
	// Write your code here
	max := 0
	m := make(map[int32]int)
	for i := 0; i < len(arr); i++ {
		m[arr[i]]++
		if m[arr[i]] > max {
			max = m[arr[i]]
		}
	}
	return int32(len(arr) - max)
}
