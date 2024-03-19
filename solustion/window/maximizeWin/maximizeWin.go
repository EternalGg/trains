package maximizeWin

func MaximizeWin(prizePositions []int, k int) int {
	biggst := prizePositions[len(prizePositions)-1]
	list := make([]int, biggst+1)
	max := 0
	for _, position := range prizePositions {
		list[position]++
	}
	for i := 0; i < len(list); i++ {
		left := i - k - 1
		if left < 0 {
			left = 0
		}
		right := i + k
		if right > len(list)-1 {
			right = len(list) - 1
		}
		now := 0
		for j := left; j <= right; j++ {
			now += list[j]
		}
		if now > max {
			max = now
		}
	}
	return max
}
