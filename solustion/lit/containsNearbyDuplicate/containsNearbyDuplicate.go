package containsNearbyDuplicate

import "math"

func containsNearbyDuplicate(nums []int, k int) bool {
	hashList := make(map[int][]int)
	for key, num := range nums {
		if len(hashList[num]) == 0 {
			hashList[num] = append(hashList[num], key)
			continue
		}
		for _, list := range hashList[num] {
			if math.Abs(float64(key-list)) <= float64(k) {
				return true
			}
		}
		hashList[num] = append(hashList[num], key)
	}
	return false
}
