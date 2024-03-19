package findNonMinOrMax

import "sort"

func findNonMinOrMax(nums []int) int {
	m, ulist := make(map[int]bool), make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] == false {
			m[nums[i]] = true
			ulist = append(ulist, nums[i])
		}
	}
	if len(ulist) <= 2 {
		return -1
	} else {
		sort.Slice(ulist, func(i, j int) bool {
			return ulist[i] < ulist[j]
		})
		return ulist[1]
	}
}
