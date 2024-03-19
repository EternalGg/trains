package longestMountain

func LongestMountain(arr []int) int {

	max, up, down := 0, 0, 0
	goingUp := false
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			if !goingUp {
				if up != 0 {
					if up+down > max {
						max = up + down
					}
					up = 0
					down = 0
				}
			}
			goingUp = true
			up++
			down = 0
		} else if arr[i] < arr[i-1] {
			if goingUp {
				goingUp = false
			}
			down++
		} else if arr[i] == arr[i-1] {
			if !goingUp {
				if up != 0 {
					if up+down > max {
						max = up + down
					}
					up = 0
					down = 0
				}
			}
			up = 0
			down = 0
			goingUp = false
		}
	}
	if !goingUp {
		if up != 0 {
			if up+down > max {
				max = up + down
			}
			up = 0
			down = 0
		}
	}
	if max < 2 {
		return 0
	} else {
		return max + 1
	}

}
