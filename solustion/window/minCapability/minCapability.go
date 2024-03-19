package minCapability

func minCapability(nums []int, k int) int {
	return 0
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	// 分成不含第一个和不含最后一个
	return max(robV1(nums[1:]), robV1(nums[:len(nums)-1]))
}

// 打家劫舍1 的代码搬过来用一下
func robV1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 补2个0在前面 便于统一处理
	nums = append([]int{0, 0}, nums...)
	for i := 2; i < len(nums); i++ {
		// 偷到第i间房子的时候的最大金额
		nums[i] = max(nums[i-2]+nums[i], nums[i-1])
	}
	return nums[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
