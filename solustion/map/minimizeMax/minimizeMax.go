package minimizeMax

import (
	"math"
	"sort"
)

func MinimizeMax(nums []int, p int) int {
	if p == 0 {
		return 0
	}
	list := []int{}
	m := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if !m[int(math.Abs(float64(nums[i]-nums[j])))] {
				m[int(math.Abs(float64(nums[i]-nums[j])))] = true
				list = append(list, int(math.Abs(float64(nums[i]-nums[j]))))
			}
		}
	}

	sort.Ints(list)
	return list[p-1]
}

func MinimizeMax2(nums []int, p int) int {
	if p == 0 {
		return 0
	}
	list := []int{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {

			list = append(list, int(math.Abs(float64(nums[i]-nums[j]))))

		}
	}

	sort.Ints(list)
	return list[p-1]
}
