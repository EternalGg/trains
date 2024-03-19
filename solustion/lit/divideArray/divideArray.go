package divideArray

import "sort"

func divideArray(nums []int, k int) [][]int {
	result := [][]int{}
	if len(nums)%3 != 0 {
		return result
	}

	sort.Ints(nums)
	for i := 0; i < len(nums)/3; i++ {
		if nums[(i+1)*3]-nums[i*3] > k {
			return [][]int{}
		}
		result = append(result, nums[i*3:(i+1)*3])
	}
	return result

	// 8 7 6 6
	// 6 6 7 8
}
