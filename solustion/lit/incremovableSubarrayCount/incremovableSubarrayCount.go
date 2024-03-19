package incremovableSubarrayCount

import "sort"

func IncremovableSubarrayCount(nums []int) int {
	count := 0
	// loop through the array from left to right
	for i := 0; i < len(nums); i++ {
		// loop through the array from right to left
		for j := len(nums) - 1; j >= i; j-- {
			// check if the subarray nums[i:j+1] is removable and increasing
			if isRemovableInc(nums, i, j) {
				// increment the count
				count++
			}
		}
	}
	// return the count
	return count
}

// isRemovableInc returns true if the subarray nums[left:right+1] is removable and increasing
func isRemovableInc(nums []int, left, right int) bool {
	// if the subarray is empty, return true
	if left > right {
		return true
	}
	// if the subarray has only one element, return true
	if left == right {
		return true
	}
	// initialize the previous element to the first element of the subarray
	prev := nums[left]
	// loop through the subarray from left + 1 to right
	for i := left + 1; i <= right; i++ {
		// if the current element is smaller than or equal to the previous element, return false
		// change <= to < to handle equal elements correctly
		if nums[i] < prev {
			return false
		}
		// update the previous element to the current element
		prev = nums[i]
	}
	// check if the subarray can be removed without breaking the strictly increasing property of the remaining array
	// if the left index is 0, or the right index is the last index, return true
	if left == 0 || right == len(nums)-1 {
		return true
	}
	// otherwise, check if the element before the left index is smaller than the element after the right index
	return nums[left-1] < nums[right+1]
}
func ZZ(nums []int) int64 {
	sort.Ints(nums)
	sum := 0
	for i := len(nums) - 1; i >= 0; i-- {
		sum += nums[i]
		if i > 0 && sum-nums[i] <= nums[i-1] {
			sum += nums[i]
		} else {
			return int64(sum)
		}
	}
	return -1
}
