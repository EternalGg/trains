package triangularSum

func TriangularSum(nums []int) int {
	for len(nums) != 1 {
		now := make([]int, len(nums)-1)
		for i := 0; i < len(nums)-1; i++ {
			value := nums[i] + nums[i+1]
			if value >= 10 {
				value = value % 10
			}
			now[i] = value
		}
		nums = now
	}
	return nums[0]
}
