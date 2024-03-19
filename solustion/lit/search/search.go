package search

func search(nums []int, target int) int {
	M := make(map[int]int, 0)
	M[target] = 0
	for _, num := range nums {
		if num == target {
			M[target]++
		}
	}
	return M[target]
}
