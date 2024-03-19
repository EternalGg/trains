package countMatchingSubarrays

func CountMatchingSubarrays(nums []int, pattern []int) int {
	result, compareList := 0, make([]int, 0)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			compareList = append(compareList, 1)
		} else if nums[i] == nums[i+1] {
			compareList = append(compareList, 0)
		} else {
			compareList = append(compareList, -1)
		}
	}
	for i := 0; i < len(compareList)-len(pattern)+1; i++ {

		if listcompare(compareList[i:i+len(pattern)], pattern) {
			result++
		}
	}
	return result
}

func listcompare(l1, l2 []int) bool {
	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}
