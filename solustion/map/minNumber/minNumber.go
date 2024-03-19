package minNumber

func minNumber(nums1 []int, nums2 []int) int {
	n1Map, n1Min := make(map[int]bool), 10
	for i := 0; i < len(nums1); i++ {
		n1Map[nums1[1]] = true
		if nums1[i] < n1Min {
			n1Min = nums1[i]
		}
	}
	n2Map, n2Min := make(map[int]bool), 10
	for i := 0; i < len(nums2); i++ {
		n2Map[nums2[1]] = true
		if nums2[i] < n2Min {
			n2Min = nums2[i]
		}
	}

	for i := 1; i < 10; i++ {
		if n1Map[i] && n2Map[i] {
			return i
		}
	}
	result := 0
	if n1Min < n2Min {
		result += (n1Min * 10) + n2Min
	} else {
		result += (n2Min * 10) + n1Min
	}
	return result
}
