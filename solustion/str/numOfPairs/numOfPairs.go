package numOfPairs

func NumOfPairs(nums []string, target string) int {
	strMap := make(map[string]int)
	for i := 0; i < len(nums); i++ {
		strMap[nums[i]]++
	}
	result := 0
	for i := 0; i < len(nums); i++ {
		length := len(nums[i])
		if length >= len(target) {
			continue
		}
		if target[:length] == nums[i] {
			cut := target[length:]
			if nums[i] == cut {
				result += strMap[cut] - 1
			} else {
				result += strMap[cut]
			}
		}
	}
	return result
}
