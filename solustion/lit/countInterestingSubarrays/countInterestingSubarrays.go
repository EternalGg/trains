package countInterestingSubarrays

func CountInterestingSubarrays(nums []int, modulo int, k int) int64 {
	list := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i]%modulo == k {
			list = append(list, 1)
		} else {
			list = append(list, 0)
		}
	}
	cntMap := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if i%modulo == k {
			cntMap[i] = true
		}
	}
	result := int64(0)
	window := 1
	for window != len(list)+1 {
		cb := false
		for i := window; i >= 0; i-- {
			if cntMap[i] {
				cb = true
				break
			}
		}
		if !cb {
			window++
			continue
		}
		for i := 0; i < len(list)-window+1; i++ {
			c := 0
			for j := i; j < window+i; j++ {
				c += list[j]
			}
			if cntMap[c] {
				result++
			}
		}
		window++
	}
	return result
}
