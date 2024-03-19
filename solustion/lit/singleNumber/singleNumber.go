package singleNumber

func singleNumber(nums []int) int {
	nMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		nMap[nums[i]]++
		if nMap[nums[i]] == 3 {
			delete(nMap, nums[i])
		}
	}
	for i, _ := range nMap {
		return i
	}
	return 0
}
