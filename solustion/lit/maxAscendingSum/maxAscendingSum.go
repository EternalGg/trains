package maxAscendingSum

func MaxAscendingSum(nums []int) int {
	Max, last, count := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < last {
			if count > Max {
				Max = count
			}
			count = nums[i]
			last = nums[i]
		} else {
			last = nums[i]
			count += nums[i]
		}
	}
	if count > Max {
		Max = count
	}
	return Max
}
