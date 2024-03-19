package findNonMinOrMax

import "sort"

func findNonMinOrMax(nums []int) int {
	m := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if !m[nums[i]] {
			m[nums[i]] = true
		}
	}
	if len(m) <= 2 {
		return -1
	} else {
		list := []int{}
		for key, _ := range m {
			list = append(list, key)
		}
		sort.Ints(list)
		return list[1]
	}
}
