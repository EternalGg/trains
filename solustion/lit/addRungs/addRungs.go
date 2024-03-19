package addRungs

func AddRungs(rungs []int, dist int) int {
	now, added := 0, 0

	for _, rung := range rungs {
		if now >= rung {
			now = rung
			continue
		} else {
			add := 0
			cut := rung - now
			if cut > 10 {
				add = Divide(now, rung, dist)
			} else {
				_, add = Add(now, rung, dist)
			}
			now = rung
			added += add
		}
	}
	return added
}

func Add(now, target, dist int) (value int, time int) {
	for true {
		if now+dist >= target {
			return now + dist, time
		} else {
			now += dist
			time++
		}
	}
	return now, 0
}

func Divide(now, target, dist int) (time int) {
	cut := target - now
	time = cut / dist

	return time
}
