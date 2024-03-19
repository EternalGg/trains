package findDifference

func findDifference(nums1 []int, nums2 []int) [][]int {
	n1map, l1 := make(map[int]bool), make([]int, 0)
	for i := 0; i < len(nums1); i++ {
		if !n1map[nums1[i]] {
			n1map[nums1[i]] = true
			l1 = append(l1, nums1[i])
		}
	}

	n2map, l2 := make(map[int]bool), make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		if !n2map[nums2[i]] {
			n2map[nums2[i]] = true
			l2 = append(l2, nums2[i])
		}
	}

	result := make([][]int, 2)
	for i := 0; i < len(l1); i++ {
		if !n2map[l1[i]] {
			result[0] = append(result[0], l1[i])
		}
	}
	for i := 0; i < len(l2); i++ {
		if !n1map[l2[i]] {
			result[1] = append(result[1], l2[i])
		}
	}
	return result
}
