package findCenter

func findCenter(edges [][]int) int {
	max, nMap := len(edges), make(map[int]int)

	for _, edge := range edges {
		for _, value := range edge {
			nMap[value]++
			if nMap[value] == max {
				return value
			}
		}
	}
	return 0
}
