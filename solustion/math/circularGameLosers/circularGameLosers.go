package circularGameLosers

func CircularGameLosers(n int, k int) []int {
	getBall, time := make(map[int]bool), 1
	getBall[1] = true
	player := 1
	for true {
		player = (time*k + player) % n
		if player == 0 {
			player = n
		}
		if getBall[player] {
			result := make([]int, 0)
			for i := 0; i < n; i++ {
				if getBall[i+1] == false {
					result = append(result, i+1)
				}
			}
			return result
		} else {
			getBall[player] = true
		}
		time++
	}
	return []int{}
}
