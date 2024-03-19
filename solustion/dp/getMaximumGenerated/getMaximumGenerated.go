package getMaximumGenerated

func GetMaximumGenerated(n int) int {
	list := make([]int, n+1)
	list[1] = 1
	for i := 2; i < len(list); i++ {
		if i%2 == 0 {
			list[i] = list[i/2]
		} else {
			list[i] = list[i/2] + list[(i/2)+1]
		}
	}
	max := list[0]
	for i := 0; i < len(list); i++ {
		if list[i] > max {
			max = list[i]
		}
	}
	return max
}
