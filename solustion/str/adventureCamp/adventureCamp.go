package adventureCamp

func AdventureCamp(expeditions []string) int {
	findList, cMap := make([]int, 0), make(map[string]bool)

	for i := 0; i < len(expeditions); i++ {
		spliteList, last, now := make([]string, 0), "", 0
		for j := 0; j < len(expeditions[i]); j++ {
			if expeditions[i] == "" {
				goto here
			} else {
				switch expeditions[i][j] {
				case '-':
					if last != "" {
						spliteList = append(spliteList, last)
						last = ""
						j++
					}
				default:
					last += string(expeditions[i][j])
				}
			}
		}
		if last != "" {
			spliteList = append(spliteList, last)
		}
		for j := 0; j < len(spliteList); j++ {
			if !cMap[spliteList[j]] {
				now++
				cMap[spliteList[j]] = true
			}
		}
		if i != 0 {
			findList = append(findList, now)
		}
	here:
	}
	max, key := 0, 1
	for i := 0; i < len(findList); i++ {
		if findList[i] > max {
			max = findList[i]
			key = i + 1
		}
	}
	if max == 0 {
		return -1
	} else {
		return key
	}

}
