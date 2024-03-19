package findDuplicates

func findDuplicates(nums []int) []int {
	bMap, result := make(map[int]bool), make([]int, 0)
	for _, num := range nums {
		if bMap[num] {
			result = append(result, num)
		} else {
			bMap[num] = true
		}
	}
	return result
}
