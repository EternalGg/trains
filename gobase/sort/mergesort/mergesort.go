package mergesort

func MergeSort(nums []int) []int {
	length := len(nums)
	left := nums[:length/2]
	right := nums[length/2:]
	return Merge(left, right)
}

func Merge(left, right []int) []int {

	return nil
}
