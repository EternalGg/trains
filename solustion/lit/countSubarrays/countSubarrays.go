package countSubarrays

func CountSubarrays(nums []int, k int) int64 {

	count := 0 // 记录满足条件的子数组数量
	n := len(nums)
	maxnum := 0
	for i := 0; i < len(nums); i++ {
		if maxnum < nums[i] {
			maxnum = nums[i]
		}
	}

	// 遍历数组中的每个数作为子数组的起始点
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			// 如果子数组中最大元素值出现次数达到要求
			if countNum(nums[i:j+1], maxnum) >= k {
				count++
			}
		}
	}

	return int64(count)
}

// 统计数组中某个元素出现的次数
func countNum(nums []int, target int) int {
	count := 0
	for _, num := range nums {
		if num == target {
			count++
		}
	}
	return count
}
