package groupThePeople

func GroupThePeople(groupSizes []int) [][]int {
	Map, list := make(map[int][]int), make([]int, 0)
	for i := 0; i < len(groupSizes); i++ {
		if len(Map[groupSizes[i]]) == 0 {
			Map[groupSizes[i]] = make([]int, 0)
			list = append(list, groupSizes[i])
		}
		Map[groupSizes[i]] = append(Map[groupSizes[i]], i)
	}

	result := make([][]int, 0)
	for _, value := range list {
		now, time := Map[value], 0
		if len(now)%value != 0 {
			time++
		}
		time += len(now) / value
		key := 0
		for i := 0; i < time; i++ {
			l := make([]int, 0)
			for j := 0; j < value; j++ {
				if key == len(now) {
					break
				}
				l = append(l, now[key])
			}
			result = append(result, l)
		}
	}
	return result
}
