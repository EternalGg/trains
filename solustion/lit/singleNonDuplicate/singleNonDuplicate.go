package singleNonDuplicate

func singleNonDuplicate(nums []int) int {
	for i := 0; i < len(nums); i += 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return 0
}
