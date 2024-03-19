package maximumSum

import "strconv"

func maximumSum(nums []int) (max int) {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if m[splitNumber(nums[i])] == 0 {
			m[splitNumber(nums[i])] = nums[i]
		} else {
			if nums[i]+m[splitNumber(nums[i])] > max {
				max = nums[i] + m[splitNumber(nums[i])]
			}
			if nums[i] > m[splitNumber(nums[i])] {
				m[splitNumber(nums[i])] = nums[i]
			}
		}
	}
	return max
}

func splitNumber(num int) int {
	result := 0
	str := strconv.Itoa(num)
	for _, i2 := range str {
		result += int(i2) - 48
	}

	return result
}
