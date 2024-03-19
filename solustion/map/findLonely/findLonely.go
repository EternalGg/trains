package findLonely

func findLonely(nums []int) []int {
	m := make(map[int]bool)
	//result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
		if m[nums[i]-1] == true {

		}
	}
	return nums
}
