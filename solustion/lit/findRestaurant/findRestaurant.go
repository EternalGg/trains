package findRestaurant

func findRestaurant(list1 []string, list2 []string) []string {
	L1Map := make(map[string]int)
	for key, str := range list1 {
		L1Map[str] = key + 1
	}
	result := make([]string, 0)
	for key, str := range list2 {
		if L1Map[str] >= 0 {
			result = append(result, str)
			L1Map[str] += key + 1
		}
	}
	if len(result) == 0 {
		return nil
	}
	minKey := L1Map[result[0]]
	for _, str := range result {
		if L1Map[str] < L1Map[result[minKey]] {
			minKey = L1Map[str]
		}
	}
	r := []string{}
	for _, value := range result {
		if L1Map[value] == minKey {
			r = append(r, value)
		}
	}

	return r
}
