package findMaxAverage

import "fmt"

func FindMaxAverage(nums []int, k int) float64 {
	if len(nums) == 1 {
		return float64(nums[0])
	}
	count, left, right, max := float64(0), 0, 0, float64(0)
	for right = 0; right < k; right++ {
		count += float64(nums[right])
	}
	max = count / float64(k)
	for right != len(nums) {
		fmt.Println(count)
		count -= float64(nums[left])
		left++
		right++
		count += float64(nums[right])
		if max < (count / float64(k)) {
			max = count / float64(k)
		}
	}

	return max
}
