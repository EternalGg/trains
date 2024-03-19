package destCity

func destCity(paths [][]string) string {
	startMap := make(map[string]bool)

	for _, path := range paths {
		startMap[path[0]] = true
	}
	for _, path := range paths {
		if !startMap[path[1]] {
			return path[1]
		}
	}
	return ""
}
