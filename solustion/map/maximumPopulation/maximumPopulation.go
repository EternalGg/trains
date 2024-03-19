package maximumPopulation

func maximumPopulation(logs [][]int) int {
	yearMap := map[int]int{}
	for i := 0; i < len(logs); i++ {
		for j := logs[i][0]; j < logs[i][1]; j++ {
			yearMap[j]++
		}
	}

	maxValue := 0
	for _, v := range yearMap {
		if v > maxValue {
			maxValue = v
		}
	}
	earlyYear := 0
	for y, v := range yearMap {
		if v == earlyYear {
			if earlyYear == 0 {
				earlyYear = y
			} else if y < earlyYear {
				earlyYear = y
			}
		}
	}
	return earlyYear
}
