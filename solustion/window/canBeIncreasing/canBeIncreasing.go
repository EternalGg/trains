package canBeIncreasing

func CanBeIncreasing(nums []int) bool {
	last, mightList, changePoint := 0, make([]int, 0), 0
	for i := 0; i < len(nums); i++ {
		if last > nums[i] {
			last = nums[i]
			changePoint++
			mightList = append(mightList, i)
			continue
		}
		last = nums[i]
	}
	if changePoint <= 1 {
		return true
	}
	return false
}
