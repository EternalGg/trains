package findSubarrays

func FindSubarrays(nums []int) bool {
	//uniqe := make(map[int]bool)
	//for _, num := range nums {
	//	if !uniqe[num] {
	//		uniqe[num] = !uniqe[num]
	//	} else {
	//		return true
	//	}
	//}

	answerMap := make(map[int]bool)
	for i := 0; i < len(nums)-1; i++ {
		if !answerMap[nums[i]+nums[i+1]] {
			answerMap[nums[i]+nums[i+1]] = true
		} else {
			return true
		}
	}
	return false
}
