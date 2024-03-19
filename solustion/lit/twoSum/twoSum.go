package twoSum

func twoSum(nums []int, target int) []int {
	Map := make(map[int]bool)
	for _, num := range nums {
		Map[num] = true

	}
	for _, num := range nums {
		if Map[target-num] == true {
			result := []int{num, target - num}
			return result
		}
	}
	return nil
}
