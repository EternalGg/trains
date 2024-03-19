package countConsistentStrings

func countConsistentStrings(allowed string, words []string) int {
	allowedMap, result := make(map[int32]bool), 0
	for _, value := range allowed {
		allowedMap[value] = true
	}
	for _, word := range words {
		for _, value := range word {
			if allowedMap[value] == false {
				break
			}
		}
		result++

	}
	return result
}
