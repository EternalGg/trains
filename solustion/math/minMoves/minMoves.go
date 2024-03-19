package minMoves

func minMoves(target int, maxDoubles int) int {
	count := 0
	for target != 1 {
		if maxDoubles == 0 {
			count += target
			return count
		}
		if target%2 == 0 && maxDoubles != 0 {
			target /= 2
			maxDoubles--
		} else {
			target--
		}
		count++
	}

	return count
}
