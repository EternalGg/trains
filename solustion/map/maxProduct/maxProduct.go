package maxProduct

func maxProduct(words []string) int {
	max := 0
	mapgroup := []map[int32]bool{}
	for i := 0; i < len(words); i++ {
		tm := make(map[int32]bool)
		for _, key := range words[i] {
			tm[key] = true
		}
		mapgroup = append(mapgroup, tm)
	}

	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			for key, _ := range mapgroup[i] {
				if mapgroup[j][key] {
					goto here
				}
			}
			if len(words[i])*len(words[j]) > max {
				max = len(words[i]) * len(words[j])
			}
		here:
		}
	}
	return max
}
