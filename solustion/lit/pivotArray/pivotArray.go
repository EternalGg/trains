package pivotArray

func pivotArray(nums []int, pivot int) []int {
	min, equal, max := make([]int, 0), 1, make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] < pivot {
			min = append(min, nums[i])
		} else if nums[i] == pivot {
			equal++
		} else {
			max = append(max, nums[i])
		}
	}
	eq := make([]int, 0)
	for i := 0; i < equal; i++ {
		eq = append(eq, pivot)
	}
	eq = append(eq, max...)
	return append(min, eq...)
}
