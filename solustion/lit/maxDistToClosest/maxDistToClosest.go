package maxDistToClosest

func MaxDistToClosest(seats []int) int {
	last, max := -1, 0
	for i := 0; i < len(seats); i++ {
		if seats[i] == 1 {
			if last == -1 {
				last = i
			} else {
				cut := i - last - 1
				now := 0
				if cut%2 == 0 {
					now = cut / 2
				} else {
					now = (cut / 2) + 1
				}
				if now > max {
					max = now
				}
				last = i
			}
		}
	}
	if last != len(seats)-1 {
		if len(seats)-1-last > max {
			max = len(seats) - 1 - last
		}
	}

	aLast := -1
	for i := len(seats) - 1; i >= 0; i-- {
		if seats[i] == 1 {
			if aLast == -1 {
				aLast = i
			} else {
				cut := aLast - i - 1
				now := 0
				if cut%2 == 0 {
					now = cut / 2
				} else {
					now = (cut / 2) + 1
				}
				if now > max {
					max = now
				}
				aLast = i
			}
		}
	}
	if aLast != 0 {
		if aLast > max {
			max = aLast
		}
	}

	return max
}
