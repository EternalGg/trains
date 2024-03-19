package countPairs

func countPairs(nums []int, k int) int {
	nMap, count := make(map[int][]int), 0

	for i, num := range nums {
		for _, key := range nMap[num] {
			if (num*key)%k == 0 {
				count++
			}
		}
		nMap[num] = append(nMap[num], i)
	}
	return count
}
