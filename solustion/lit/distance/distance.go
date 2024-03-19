package distance

import "math"

func distance(nums []int) []int64 {
	result := []int64{}
	m := map[int][]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = append(m[nums[i]], i)
	}
	for i := 0; i < len(nums); i++ {
		now := int64(0)
		for _, i3 := range m[nums[i]] {
			now += int64(math.Abs(float64(i - i3)))
		}
		result = append(result, now)
	}
	return result
}
