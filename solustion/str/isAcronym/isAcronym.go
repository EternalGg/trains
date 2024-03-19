package isAcronym

func isAcronym(words []string, s string) bool {
	if len(s) != len(words) {
		return false
	}
	check := ""
	for i := 0; i < len(words); i++ {
		check += string(words[i][0])
	}
	if check == s {
		return true
	} else {
		return false
	}
}

func countPairs(nums []int, target int) int {
	result := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] < target {
				result++
			}
		}
	}
	return result
}
