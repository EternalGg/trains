package distributeCandiesi

func DistributeCandies(candyType []int) int {
	typeList, candieMap := make([]int, 0), make(map[int]int)

	for _, value := range candyType {
		if candieMap[value] == 0 {
			typeList = append(typeList, value)
		}
		candieMap[value]++
	}
	if len(candyType)/2 >= len(typeList) {
		return len(candyType) / 2
	} else {
		return len(typeList)
	}

}
