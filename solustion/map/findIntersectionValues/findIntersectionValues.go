package findIntersectionValues

func findIntersectionValues(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]] = true
	}
	result := []int{0, 0}
	m2 := make(map[int]bool)
	for i := 0; i < len(nums2); i++ {
		m2[nums2[i]] = true
		if m[nums2[i]] {
			result[1]++
		}
	}
	for i := 0; i < len(nums1); i++ {
		if m2[nums1[i]] {
			result[0]++
		}
	}

	return result
}
