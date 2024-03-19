package numberOfPoints

func numberOfPoints(nums [][]int) int {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		for j := nums[i][0]; j < nums[i][1]; j++ {
			if m[j] == false {
				m[nums[i][j]] = true
			}
		}
	}
	return len(m)
}
