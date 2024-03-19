package maximumDifference

func MaximumDifference(nums []int) int {
	max := -1
	for i := len(nums) - 1; i > 0; i-- {
		for j := i - 1; j > 0; j-- {
			if nums[i]-nums[j] > max && nums[i] != nums[j] {
				max = nums[i] - nums[j]
			}
		}
	}
	return max
}
