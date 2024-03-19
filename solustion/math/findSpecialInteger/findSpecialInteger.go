package findSpecialInteger

func FindSpecialInteger(arr []int) int {
	countMap := make(map[int]int)
	length := len(arr) / 4
	for i := 0; i < len(arr); i++ {
		countMap[arr[i]]++
		if countMap[arr[i]] >= length {
			return arr[i]
		}
	}
	return 0
}
