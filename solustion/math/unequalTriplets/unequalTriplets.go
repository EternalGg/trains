package unequalTriplets

func UnequalTriplets(nums []int) int {
	count, nMap := 0, make(map[int][]int)

	for i, num := range nums {
		nMap[num] = append(nMap[num], i)
	}

	return count
}
