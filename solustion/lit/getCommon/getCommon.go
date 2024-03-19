package getCommon

func GetCommon(nums1 []int, nums2 []int) int {
	nMap := make(map[int]bool)
	for i := 0; i < len(nums1); i++ {
		nMap[nums1[i]] = true
	}
	for i := 0; i < len(nums2); i++ {
		if nMap[nums2[i]] == true {
			return nums2[i]
		}
	}
	return -1
}
