package triangleType

func triangleType(nums []int) string {
	if len(nums) > 3 {
		return "none"
	}
	if nums[0] == nums[1] && nums[1] == nums[2] {
		return "equilateral"
	} else if nums[0] == nums[1] || nums[1] == nums[2] || nums[0] == nums[2] {
		return "isosceles"
	} else {
		return "scalene"
	}
}
